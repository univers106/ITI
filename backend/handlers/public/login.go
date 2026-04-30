package public

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/univers106/ITI/middlewares/database_middleware"
	"github.com/univers106/ITI/middlewares/sessions_middleware"
)

func PostLogin(c *echo.Context) error {
	sessionStorage, err := sessions_middleware.GetSessionStorage(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to get session")
	}

	_, err = sessions_middleware.GetKeyFromCookies(c)
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

	db, err := database_middleware.GetDatabase(c)
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

	sessions_middleware.SetKeyToCookies(c, sessionKey)

	return c.JSON(http.StatusOK, map[string]string{"message": "ok"})
}
