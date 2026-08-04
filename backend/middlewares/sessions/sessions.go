package sessions

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/univers106/ITI/database"
	user_database_middleware "github.com/univers106/ITI/middlewares/databases/user"
)

const AuthSession = "auth"

var (
	ErrUnauthorized = errors.New("unauthorized")
	ErrBadCookies   = errors.New("bad cookies")
)

const (
	SessionTimeout     = 3 * time.Hour
	SessionIdleTimeout = 10 * time.Minute
)

func NewSessionsMiddleware(store SessionStorage) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			c.Set("_session_storage", store)

			return next(c)
		}
	}
}

func getSessionStorage(c *echo.Context) (SessionStorage, error) {
	store, err := echo.ContextGet[SessionStorage](c, "_session_storage")
	if err != nil {
		return nil, fmt.Errorf("failed to get session store: %w", err)
	}

	return store, nil
}

func getKeyFromCookies(c *echo.Context) (string, error) {
	cookie, err := c.Cookie("session_key")
	if err != nil {
		return "", ErrUnauthorized
	}

	return cookie.Value, nil
}

func setKeyToCookies(c *echo.Context, sessionKey string) {
	cookie := new(http.Cookie)
	cookie.Name = "session_key"
	cookie.Value = sessionKey
	cookie.Path = "/"
	cookie.HttpOnly = true
	cookie.Secure = false
	cookie.SameSite = http.SameSiteNoneMode
	cookie.MaxAge = int((time.Hour + SessionIdleTimeout).Seconds())
	c.SetCookie(cookie)
}

func deleteKeyFromCookies(c *echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = "session_key"
	cookie.Value = ""
	cookie.Path = "/"
	cookie.HttpOnly = true
	cookie.Secure = false
	cookie.SameSite = http.SameSiteNoneMode
	cookie.MaxAge = -1
	cookie.Expires = time.Unix(1, 0)
	c.SetCookie(cookie)
}

func getUserFromSession(c *echo.Context) (*database.User, error) {
	sessionStorage, err := getSessionStorage(c)
	if err != nil {
		return nil, echo.ErrInternalServerError
	}

	sessionKey, err := getKeyFromCookies(c)
	if err != nil {
		return nil, echo.ErrUnauthorized
	}

	userLogin, err := sessionStorage.GetLoginFromSession(sessionKey)
	if err != nil {
		deleteKeyFromCookies(c)

		return nil, echo.ErrBadRequest
	}

	db, err := user_database_middleware.Get(c)
	if err != nil {
		return nil, echo.ErrInternalServerError
	}

	user, err := db.GetByLogin(userLogin)
	if err != nil {
		deleteKeyFromCookies(c)

		return nil, echo.ErrInternalServerError
	}

	return user, nil
}
