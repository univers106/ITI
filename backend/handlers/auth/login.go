package auth

import (
	"net/http"

	"github.com/labstack/echo/v5"
	sessionsMiddleware "github.com/univers106/ITI/middlewares/sessions"
)

func PostLogin(c *echo.Context) error {
	session, err := sessionsMiddleware.GetUserSession(c)
	if err != nil {
		return err
	}
	username := c.FormValue("username")
	if username == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "bad request"})
	}
	password := c.FormValue("password")
	if password == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "bad request"})
	}
	if username != "test" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "unauthorized"})
	}
	if password != "pass" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "unauthorized"})
	}

	session.Values["username"] = username
	if err := session.Save(c.Request(), c.Response()); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "ok"})
}
