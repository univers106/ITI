package database

import "slices"

type User struct {
	Login       string   `json:"login"`
	Name        string   `json:"name"`
	Permissions []string `json:"permissions"`
	newPassword string
}

func (u *User) HasPermission(permission string) bool {
	if slices.Contains(u.Permissions, PermSuperUser) {
		return true
	}

	return slices.Contains(u.Permissions, permission)
}

func (u *User) SetPassword(newPassword string) {
	u.newPassword = newPassword
}

func (u *User) IsPasswordUpdated() (bool, string) {
	return u.newPassword != "", u.newPassword
}
