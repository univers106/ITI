package auth

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/univers106/ITI/middlewares/sessions"
)

func GetMe(c *echo.Context) error {
	user, err := sessions.GetUser(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}
