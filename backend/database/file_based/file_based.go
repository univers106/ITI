package filebased

import (
	"slices"
	"sync"

	"github.com/univers106/ITI/database"
)

type FileBasedDatabase struct {
	users   map[int]safeUser
	mu      sync.RWMutex
	dirPath string
}

func (f *FileBasedDatabase) GetUserByID(id int) (*database.User, error) {
	f.mu.RLock()
	defer f.mu.RUnlock()

	user, ok := f.users[id]
	if !ok {
		return nil, database.ErrUserNotFound
	}

	return user.GetUser(), nil
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

func (f *FileBasedDatabase) UserAuthentication(
	login string,
	password string,
) (*database.User, error) {
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

	_, err := f.GetUserByLogin(login)
	if err == nil {
		return database.ErrUserExists
	}

	f.mu.Lock()
	defer f.mu.Unlock()

	newUser := safeUser{User: database.User{Login: login, Name: name}}
	newUser.SetPassword(password)

	for id := range f.users {
		if id >= newUser.ID {
			newUser.ID = id + 1
		}
	}

	f.users[newUser.ID] = newUser

	return nil
}

func (f *FileBasedDatabase) DeleteUser(userId int) error {
	defer f.save()

	f.mu.Lock()
	defer f.mu.Unlock()

	if _, ok := f.users[userId]; !ok {
		return database.ErrUserNotFound
	}

	delete(f.users, userId)

	return nil
}

func (f *FileBasedDatabase) ChangeUserPassword(user_id int, password string) error {
	defer f.save()

	f.mu.Lock()
	defer f.mu.Unlock()

	user, ok := f.users[user_id]
	if !ok {
		return database.ErrUserNotFound
	}

	user.SetPassword(password)
	f.users[user_id] = user

	return nil
}

func (f *FileBasedDatabase) ChangeUserLogin(user_id int, login string) error {
	defer f.save()

	f.mu.Lock()
	defer f.mu.Unlock()

	user, ok := f.users[user_id]
	if !ok {
		return database.ErrUserNotFound
	}

	user.Login = login
	f.users[user_id] = user

	return nil
}

func (f *FileBasedDatabase) ChangeUserName(user_id int, name string) error {
	defer f.save()

	f.mu.Lock()
	defer f.mu.Unlock()

	user, ok := f.users[user_id]
	if !ok {
		return database.ErrUserNotFound
	}

	user.Name = name
	f.users[user_id] = user

	return nil
}

func (f *FileBasedDatabase) UserAddPermissions(user_id int, permission string) error {
	defer f.save()

	f.mu.Lock()
	defer f.mu.Unlock()

	user, ok := f.users[user_id]
	if !ok {
		return database.ErrUserNotFound
	}

	if slices.Contains(user.Permissions, permission) {
		return database.ErrAlreadyExists
	}

	user.Permissions = append(user.Permissions, permission)
	f.users[user_id] = user

	return nil
}

func (f *FileBasedDatabase) UserRemovePermissions(user_id int, permission string) error {
	defer f.save()

	f.mu.Lock()
	defer f.mu.Unlock()

	user, ok := f.users[user_id]
	if !ok {
		return database.ErrUserNotFound
	}

	if !slices.Contains(user.Permissions, permission) {
		return database.ErrPermissionNotFound
	}

	user.Permissions = slices.DeleteFunc(user.Permissions, func(p string) bool {
		return p == permission
	})

	f.users[user_id] = user

	return nil
}

func (f *FileBasedDatabase) UserCheckPermission(user_id int, permission string) (bool, error) {
	f.mu.Lock()
	defer f.mu.Unlock()

	user, ok := f.users[user_id]
	if !ok {
		return false, database.ErrUserNotFound
	}

	return slices.Contains(user.Permissions, permission), nil
}

func NewFileBasedDatabase(dir string) *FileBasedDatabase {
	newUser := FileBasedDatabase{dirPath: dir}
	createDirIfNotExists(dir)

	var err error

	newUser.users, err = loadStructFromJsonFile[map[int]safeUser](dir + "/users.json")
	if err != nil {
		newUser.users = map[int]safeUser{}
	}

	newUser.save()

	return &newUser
}

func (f *FileBasedDatabase) save() {
	f.mu.Lock()
	defer f.mu.Unlock()

	createDirIfNotExists(f.dirPath)
	saveStructToJsonFile(f.dirPath+"/users.json", f.users)
}
