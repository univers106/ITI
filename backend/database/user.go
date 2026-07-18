package database

import "slices"

type User struct {
	Login       string   `json:"login"`
	Name        string   `json:"name"`
	Permissions []string `json:"permissions"`
}

func (u *User) HasPermission(permission string) bool {
	if slices.Contains(u.Permissions, PermSuperUser) {
		return true
	}

	return slices.Contains(u.Permissions, permission)
}
