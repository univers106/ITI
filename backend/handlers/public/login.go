package public

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/univers106/ITI/middlewares/databaseMiddleware"
	"github.com/univers106/ITI/middlewares/sessionsMiddleware"
)

func PostLogin(c *echo.Context) error {
	sessionStorage, err := sessionsMiddleware.GetSessionStorage(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to get session")
	}

	_, err = sessionsMiddleware.GetKeyFromCookies(c)
	if err == nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			"session already exists, try logging out first",
		)
	}

	loginValue := c.FormValue("login")
	passwordValue := c.FormValue("password")

	if passwordValue == "" || loginValue == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "login or password value is null")
	}

	db, err := databaseMiddleware.GetDatabase(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to get database")
	}

	user, err := db.UserAuthentication(loginValue, passwordValue)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
	}

	sessionKey, err := sessionStorage.NewSession(user.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to create session")
	}

	sessionsMiddleware.SetKeyToCookies(c, sessionKey)

	return c.JSON(http.StatusOK, map[string]string{"message": "ok"})
}
