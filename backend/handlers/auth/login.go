package auth

import (
	"net/http"

	"github.com/labstack/echo/v5"
	databaseMiddleware "github.com/univers106/ITI/middlewares/database"
	"github.com/univers106/ITI/middlewares/sessionsMiddleware"
)

func PostLogin(c *echo.Context) error {
	session, err := sessionsMiddleware.GetAuthSessionFromContext(c)
	if err != nil {
		return err
	}
	loginValue := c.FormValue("login")
	if loginValue == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "bad request"})
	}
	passwordValue := c.FormValue("password")
	if passwordValue == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "bad request"})
	}

	db, err := databaseMiddleware.GetDatabase(c)
	if err != nil {
		return err
	}

	user, err := db.UserAuthentication(loginValue, passwordValue)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "unauthorized"})
	}

	session.Values["login"] = user.Login
	if err := session.Save(c.Request(), c.Response()); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "ok"})
}
