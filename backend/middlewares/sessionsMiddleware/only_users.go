package sessionsMiddleware

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/univers106/ITI/database"
)

func OnlyUsersMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) error {
		user, err := GetUserFromSessionContext(c)
		if err != nil {
			if errors.Is(err, ErrUnauthorized) {
				return c.JSON(http.StatusUnauthorized, map[string]string{"message": "unauthorized. This endpoint requires authentication."})
			}
			return err
		}
		c.Set("user", user)
		return next(c)
	}
}

func GetUserFromContext(c *echo.Context) (database.User, error) {
	return echo.ContextGet[database.User](c, "user")
}
