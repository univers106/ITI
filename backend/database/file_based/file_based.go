package filebased

import (
	"slices"
	"sync"

	"github.com/univers106/ITI/database"
)

type FileBasedDatabase struct {
	users   []safeUser
	mu      sync.RWMutex
	dirPath string
}

func (f *FileBasedDatabase) GetUserByID(id int) (*database.User, error) {
	f.mu.RLock()
	defer f.mu.RUnlock()
	for _, user := range f.users {
		if user.ID == id {
			return user.GetUser(), nil
		}
	}
	return nil, database.ErrUserNotFound
}
func (f *FileBasedDatabase) GetUserByLogin(login string) (*database.User, error) {
	f.mu.RLock()
	defer f.mu.RUnlock()
	for _, user := range f.users {
		if user.Login == login {
			return user.GetUser(), nil
		}
	}
	return nil, database.ErrUserNotFound
}
func (f *FileBasedDatabase) UserAuthentication(login string, password string) (*database.User, error) {
	f.mu.RLock()
	defer f.mu.RUnlock()
	for _, user := range f.users {
		if user.Login == login {
			if user.CheckPassword(password) {
				return user.GetUser(), nil
			}
			return nil, database.ErrIncorrectPassword
		}
	}
	return nil, database.ErrUserNotFound
}
func (f *FileBasedDatabase) CreateUser(login string, name string, password string) error {
	defer f.save()

	if password == "" {
		return database.ErrPasswordEmpty
	}
	if login == "" {
		return database.ErrLoginEmpty
	}
	if name == "" {
		return database.ErrNameEmpty
	}
	if _, err := f.GetUserByLogin(login); err == nil {
		return database.ErrUserExists
	}
	f.mu.Lock()
	defer f.mu.Unlock()
	newUser := safeUser{User: database.User{Login: login, Name: name}}
	newUser.SetPassword(password)

	newUser.ID = len(f.users) + 1
	f.users = append(f.users, newUser)
	return nil
}
func (f *FileBasedDatabase) DeleteUser(id int) error {
	defer f.save()

	f.mu.Lock()
	defer f.mu.Unlock()
	for i, user := range f.users {
		if user.ID == id {
			f.users = append(f.users[:i], f.users[i+1:]...)
			return nil
		}
	}
	return database.ErrUserNotFound
}
func (f *FileBasedDatabase) ChangeUserPassword(user database.User, password string) error {
	defer f.save()

	f.mu.Lock()
	defer f.mu.Unlock()
	for i, u := range f.users {
		if u.ID == user.ID {
			f.users[i].SetPassword(password)
			return nil
		}
	}
	return database.ErrUserNotFound
}

func (f *FileBasedDatabase) UserAddPermissions(user_id int, permission string) error {
	defer f.save()

	f.mu.Lock()
	defer f.mu.Unlock()
	for i, _ := range f.users {
		user := &f.users[i]
		if user.ID == user_id {
			if slices.Contains(user.Permissions, permission) {
				return database.ErrAlreadyExists
			}
			user.Permissions = append(user.Permissions, permission)
			return nil
		}
	}
	return database.ErrUserNotFound
}
func (f *FileBasedDatabase) UserRemovePermissions(user_id int, permission string) error {
	defer f.save()

	f.mu.Lock()
	defer f.mu.Unlock()
	for i, _ := range f.users {
		user := &f.users[i]
		if user.ID == user_id {
			for i, p := range user.Permissions {
				if p == permission {
					user.Permissions = append(user.Permissions[:i], user.Permissions[i+1:]...)
					return nil
				}
			}
			return database.ErrPermissionNotFound
		}
	}
	return database.ErrUserNotFound
}

func (f *FileBasedDatabase) UserCheckPermission(user_id int, permission string) (bool, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	for _, user := range f.users {
		if user.ID == user_id {
			if slices.Contains(user.Permissions, permission) {
				return true, nil
			}
			return false, nil
		}
	}
	return false, database.ErrUserNotFound
}

func NewFileBasedDatabase(dir string) *FileBasedDatabase {
	new := FileBasedDatabase{dirPath: dir}
	createDirIfNotExists(dir)
	var err error
	new.users, err = loadStructFromJsonFile[[]safeUser](dir + "/users.json")
	if err != nil {
		new.users = []safeUser{}
	}
	new.save()
	return &new
}

func (f *FileBasedDatabase) save() {
	f.mu.Lock()
	defer f.mu.Unlock()
	createDirIfNotExists(f.dirPath)
	saveStructToJsonFile(f.dirPath+"/users.json", f.users)
}
