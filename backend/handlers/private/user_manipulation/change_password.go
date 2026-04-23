package userManipulation

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v5"
	"github.com/univers106/ITI/database"
	"github.com/univers106/ITI/middlewares/sessionsMiddleware"
)

func PostChangePassword(c *echo.Context) error {
	_, db, httpErr := sessionsMiddleware.GetUserDbCheckPermision(
		c,
		database.PermUsersManipulation,
	)
	if httpErr != nil {
		return httpErr
	}

	userIdValue := c.FormValue("userId")
	passwordValue := c.FormValue("password")

	if passwordValue == "" || userIdValue == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Some value is null")
	}

	userId, err := strconv.Atoi(userIdValue)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid userId")
	}

	err = db.ChangeUserPassword(userId, passwordValue)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			"failed to change password: "+err.Error(),
		)
	}

	return c.NoContent(http.StatusOK)
}
