package user_manipulation

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v5"
	"github.com/univers106/ITI/database"
	"github.com/univers106/ITI/middlewares/sessions_middleware"
)

func PostDelete(c *echo.Context) error {
	_, db, httpErr := sessions_middleware.GetUserDbCheckPermision(
		c,
		database.PermUsersManipulation,
	)
	if httpErr != nil {
		return httpErr
	}

	userIdValue := c.FormValue("userId")
	if userIdValue == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "userId is null")
	}

	userId, err := strconv.Atoi(userIdValue)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid userId")
	}

	sessionStorage, err := sessions_middleware.GetSessionStorage(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to get session storage")
	}

	err = db.DeleteUser(userId)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			"failed to delete user: "+err.Error(),
		)
	}

	err = sessionStorage.DeleteUserSessions(userId)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			"failed to delete user sessions"+err.Error(),
		)
	}

	return c.NoContent(http.StatusOK)
}
