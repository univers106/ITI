package public

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/univers106/ITI/middlewares/databaseMiddleware"
	"github.com/univers106/ITI/middlewares/sessionsMiddleware"
)

func PostLogin(c *echo.Context) error {
	session, err := sessionsMiddleware.GetAuthSessionFromContext(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to get session")
	}

	if !session.IsNew {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "you already logged in"})
	}

	loginValue := c.FormValue("login")
	passwordValue := c.FormValue("password")

	if passwordValue == "" || loginValue == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "login or password value is null")
	}

	db, err := databaseMiddleware.GetDatabase(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to get database")
	}

	user, err := db.UserAuthentication(loginValue, passwordValue)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
	}

	session.Values["login"] = user.Login

	err = session.Save(c.Request(), c.Response())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to save session")
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "ok"})
}
