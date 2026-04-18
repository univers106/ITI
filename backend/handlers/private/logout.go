package private

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/univers106/ITI/middlewares/sessionsMiddleware"
)

func PostLogut(c *echo.Context) error {
	session, err := sessionsMiddleware.GetAuthSessionFromContext(c)
	if err != nil {
		return err
	}
	if session.IsNew {
		return c.JSON(http.StatusOK, map[string]string{"message": "ok, but why?"})
	}
	session.Options.MaxAge = -1
	session.Save(c.Request(), c.Response())
	return c.JSON(http.StatusOK, map[string]string{"message": "ok"})
}
