package public

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/univers106/ITI/middlewares/databaseMiddleware"
	"github.com/univers106/ITI/middlewares/sessionsMiddleware"
)

func PostLogin(c *echo.Context) error {
	session, err := sessionsMiddleware.GetAuthSessionFromContext(c)
	if err != nil {
		return fmt.Errorf("failed to get auth session: %w", err)
	}

	if !session.IsNew {
		return c.JSON(http.StatusOK, map[string]string{"message": "ok, but you already logged in"})
	}

	loginValue := c.FormValue("login")
	if loginValue == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "login value is null")
	}

	passwordValue := c.FormValue("password")
	if passwordValue == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "password value is null")
	}

	db, err := databaseMiddleware.GetDatabase(c)
	if err != nil {
		return fmt.Errorf("failed to get database: %w", err)
	}

	user, err := db.UserAuthentication(loginValue, passwordValue)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
	}

	session.Values["login"] = user.Login

	err = session.Save(c.Request(), c.Response())
	if err != nil {
		return fmt.Errorf("failed to save session: %w", err)
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "ok"})
}
