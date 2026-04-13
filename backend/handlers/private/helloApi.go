package private

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/univers106/ITI/middlewares/sessions"
)

func GetHello(c *echo.Context) error {
	session, err := sessionsMiddleware.GetUserSession(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "unauthorized")
	}
	username := session.Values["username"]
	if username == nil {
		return c.JSON(http.StatusUnauthorized, "unauthorized")
	}
	usernameStr := username.(string)
	return c.JSON(http.StatusOK, "Hello "+usernameStr)
}
