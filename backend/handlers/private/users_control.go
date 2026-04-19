package private

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/univers106/ITI/database"
	"github.com/univers106/ITI/middlewares/databaseMiddleware"
	"github.com/univers106/ITI/middlewares/sessionsMiddleware"
)

func PostCreateUser(c *echo.Context) error {
	user, err := sessionsMiddleware.GetUserFromContext(c)
	if err != nil {
		return fmt.Errorf("failed to get user from context: %w", err)
	}

	if !user.HasPermission(database.PermUsersManipulation) {
		return c.JSON(http.StatusForbidden, "You do not have permission to manipulate users")
	}

	db, err := databaseMiddleware.GetDatabase(c)
	if err != nil {
		return fmt.Errorf("failed to get database: %w", err)
	}

	loginValue := c.FormValue("login")
	if loginValue == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "login value is null")
	}

	passwordValue := c.FormValue("password")
	if passwordValue == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "password value is null")
	}

	nameValue := c.FormValue("name")
	if nameValue == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "name value is null")
	}

	err = db.CreateUser(loginValue, nameValue, passwordValue)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return c.JSON(http.StatusOK, "User created successfully")
}
