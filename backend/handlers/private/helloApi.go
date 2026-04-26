package private

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/univers106/ITI/middlewares/sessionsMiddleware"
)

func GetHello(c *echo.Context) error {
	user, err := sessionsMiddleware.GetUser(c)
	if err != nil {
		return fmt.Errorf("failed to get user from context: %w", err)
	}

	return c.JSON(http.StatusOK, "Hello "+user.Name+", you login is "+user.Login)
}
