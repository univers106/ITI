package sessions_middleware

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/univers106/ITI/database"
	"github.com/univers106/ITI/middlewares/database_middleware"
)

func OnlyUsersMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) error {
		user, err := GetUserFromSession(c)
		if err != nil {
			if errors.Is(err, ErrUnauthorized) {
				return echo.NewHTTPError(
					http.StatusUnauthorized,
					"unauthorized. This endpoint requires authentication.",
				)
			}

			return err
		}

		c.Set("user", user)

		return next(c)
	}
}

func GetUser(c *echo.Context) (*database.User, error) {
	return echo.ContextGet[*database.User](c, "user")
}

func GetUserDbCheckPermision(
	c *echo.Context,
	permission string,
) (*database.User, database.Database, *echo.HTTPError) {
	user, err := GetUser(c)
	if err != nil {
		return nil, nil, echo.NewHTTPError(
			http.StatusInternalServerError,
			"failed to get user from context",
		)
	}

	if !user.HasPermission(permission) {
		return nil, nil, echo.NewHTTPError(
			http.StatusForbidden,
			"You do not have permission to manipulate users",
		)
	}

	db, err := database_middleware.GetDatabase(c)
	if err != nil {
		return nil, nil, echo.NewHTTPError(http.StatusInternalServerError, "failed to get database")
	}

	return user, db, nil
}
