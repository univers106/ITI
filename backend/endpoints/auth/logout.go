package auth

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/univers106/ITI/middlewares/sessions"
)

func PostLogout(c *echo.Context) error {
	err := sessions.DeleteSession(c)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "ok"})
}
