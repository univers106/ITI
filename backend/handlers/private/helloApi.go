package private

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/univers106/ITI/middlewares/sessions_middleware"
)

func GetHello(c *echo.Context) error {
	user, err := sessions_middleware.GetUser(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to get user from context")
	}

	return c.JSON(http.StatusOK, "Hello "+user.Name+", you login is "+user.Login)
}
