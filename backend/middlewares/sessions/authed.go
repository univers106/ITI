package sessions

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/univers106/ITI/database"
)

func Authed(permissions []string) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			user, err := getUserFromSession(c)
			if err != nil {
				return err
			}

			if !user.HasPermission(database.PermSuperUser) {
				for _, permission := range permissions {
					if !user.HasPermission(permission) {
						return echo.NewHTTPError(
							http.StatusForbidden,
							"forbidden. You do not have the required permissions.",
						)
					}
				}
			}

			c.Set("user", user)

			return next(c)
		}
	}
}
