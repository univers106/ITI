package database

import "slices"

type User struct {
	Name        string   `json:"name"`
	Login       string   `json:"login"`
	Permissions []string `json:"permissions"`
}

func (u *User) HasPermission(permission string) bool {
	if slices.Contains(u.Permissions, PermSuperUser) {
		return true
	}

	return slices.Contains(u.Permissions, permission)
}
