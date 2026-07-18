package user_manipulation

import (
	"net/http"

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

	userIdValue := c.FormValue("userLogin")
	if userIdValue == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "userLogin is null")
	}

	user, err := db.GetByLogin(userIdValue)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid userLogin")
	}

	sessionStorage, err := sessions_middleware.GetSessionStorage(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to get session storage")
	}

	err = db.DeleteUser(user.Login)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			"failed to delete user: "+err.Error(),
		)
	}

	err = sessionStorage.DeleteUserSessions(user.Login)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			"failed to delete user sessions"+err.Error(),
		)
	}

	return c.NoContent(http.StatusOK)
}
