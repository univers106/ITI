package private

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/univers106/ITI/database"
	"github.com/univers106/ITI/middlewares/databaseMiddleware"
	"github.com/univers106/ITI/middlewares/sessionsMiddleware"
)

func PostCreateUser(c *echo.Context) error {
	user, err := sessionsMiddleware.GetUserFromContext(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to get user from context")
	}

	if !user.HasPermission(database.PermUsersManipulation) {
		return c.JSON(http.StatusForbidden, "You do not have permission to manipulate users")
	}

	db, err := databaseMiddleware.GetDatabase(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to get database")
	}

	loginValue := c.FormValue("login")
	passwordValue := c.FormValue("password")
	nameValue := c.FormValue("name")

	if loginValue == "" || nameValue == "" || passwordValue == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "some value is null")
	}

	err = db.CreateUser(loginValue, nameValue, passwordValue)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to create user")
	}

	return c.JSON(http.StatusOK, "User created successfully")
}
