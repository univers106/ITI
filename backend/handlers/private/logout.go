package private

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/univers106/ITI/middlewares/sessionsMiddleware"
)

func PostLogout(c *echo.Context) error {
	session, err := sessionsMiddleware.GetAuthSessionFromContext(c)
	if err != nil {
		return fmt.Errorf("failed to get auth session from context: %w", err)
	}

	if session.IsNew {
		return c.JSON(http.StatusOK, map[string]string{"message": "ok, but why?"})
	}

	session.Options.MaxAge = -1

	err = session.Save(c.Request(), c.Response())
	if err != nil {
		return fmt.Errorf("failed to save auth session: %w", err)
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "ok"})
}
