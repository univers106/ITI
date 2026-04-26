package userManipulation

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/univers106/ITI/database"
	"github.com/univers106/ITI/middlewares/sessionsMiddleware"
)

func PostCreate(c *echo.Context) error {
	_, db, httpErr := sessionsMiddleware.GetUserDbCheckPermision(
		c,
		database.PermUsersManipulation,
	)
	if httpErr != nil {
		return httpErr
	}

	loginValue := c.FormValue("login")
	passwordValue := c.FormValue("password")
	nameValue := c.FormValue("name")

	if loginValue == "" || nameValue == "" || passwordValue == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "some value is null")
	}

	err := db.CreateUser(loginValue, nameValue, passwordValue)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			"failed to create user: "+err.Error(),
		)
	}

	return c.NoContent(http.StatusOK)
}
