package public

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

func GetHello(c *echo.Context) error {
	return c.JSON(http.StatusOK, "Hello")
}
