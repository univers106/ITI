package sessions

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/univers106/ITI/database"
)

var (
	ErrFailedToCreateSession = errors.New("failed to create session")
)

func NewSession(c *echo.Context, login string) error {
	sessionStorage, err := getSessionStorage(c)
	if err != nil {
		return err
	}

	sessionKey, err := sessionStorage.NewSession(login)
	if err != nil {
		return err
	}

	setKeyToCookies(c, sessionKey)
	return nil
}

func DeleteSession(c *echo.Context) error {
	sessionStorage, err := getSessionStorage(c)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			"failed to get auth session from context: "+err.Error(),
		)
	}

	sessionKey, err := getKeyFromCookies(c)
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

	deleteKeyFromCookies(c)

	return nil
}

func GetUser(c *echo.Context) (*database.User, error) {
	return echo.ContextGet[*database.User](c, "user")
}
