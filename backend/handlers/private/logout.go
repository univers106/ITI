package private

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/univers106/ITI/middlewares/sessionsMiddleware"
)

func PostLogout(c *echo.Context) error {
	sessionStorage, err := sessionsMiddleware.GetSessionStorage(c)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			"failed to get auth session from context: "+err.Error(),
		)
	}

	sessionKey, err := sessionsMiddleware.GetKeyFromCookies(c)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			"failed to get session key from cookies: "+err.Error(),
		)
	}

	err = sessionStorage.DeleteSession(sessionKey)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			"failed to delete session: "+err.Error(),
		)
	}

	sessionsMiddleware.DeleteCookies(c)

	return c.JSON(http.StatusOK, map[string]string{"message": "ok"})
}
