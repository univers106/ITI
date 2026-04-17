package private

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/univers106/ITI/middlewares/sessionsMiddleware"
)

func GetHello(c *echo.Context) error {
	user, err := sessionsMiddleware.GetUserFromContext(c)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, "Hello "+user.Name+", you login is "+user.Login)
}
