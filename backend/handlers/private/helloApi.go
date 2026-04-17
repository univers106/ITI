package private

import (
	"net/http"

	"github.com/labstack/echo/v5"
	databaseMiddleware "github.com/univers106/ITI/middlewares/database"
	sessionsMiddleware "github.com/univers106/ITI/middlewares/sessions"
)

func GetHello(c *echo.Context) error {
	session, err := sessionsMiddleware.GetUserSession(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "unauthorized")
	}
	login := session.Values["login"]
	if login == nil {
		return c.JSON(http.StatusUnauthorized, "unauthorized")
	}
	loginStr := login.(string)
	db, err := databaseMiddleware.GetDatabase(c)
	if err != nil {
		return err
	}
	user, err := db.GetUserByLogin(loginStr)
	if err != nil {
		return c.JSON(http.StatusNotFound, "user not found")
	}
	return c.JSON(http.StatusOK, "Hello "+user.Name+", you login is "+user.Login)
}
