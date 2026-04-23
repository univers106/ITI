package sessionsMiddleware

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/univers106/ITI/database"
	"github.com/univers106/ITI/middlewares/databaseMiddleware"
)

const AuthSession = "auth"

var (
	ErrUnauthorized = errors.New("unauthorized")
	ErrBadCookies   = errors.New("bad cookies")
)

func NewSessionsMiddleware(store SessionStorage) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			c.Set("_session_storage", store)

			return next(c)
		}
	}
}

func GetSessionStorage(c *echo.Context) (SessionStorage, error) {
	store, err := echo.ContextGet[SessionStorage](c, "_session_storage")
	if err != nil {
		return nil, fmt.Errorf("failed to get session store: %w", err)
	}

	return store, nil
}

func GetKeyFromCookies(c *echo.Context) (string, error) {
	cookie, err := c.Cookie("session_key")
	if err != nil {
		return "", ErrUnauthorized
	}

	return cookie.Value, nil
}

func SetKeyToCookies(c *echo.Context, sessionKey string) {
	cookie := new(http.Cookie)
	cookie.Name = "session_key"
	cookie.Value = sessionKey
	cookie.Path = "/"
	cookie.HttpOnly = true
	cookie.Secure = true
	cookie.SameSite = http.SameSiteStrictMode
	cookie.MaxAge = int((time.Hour + sessionIdleTimeout).Seconds())
	c.SetCookie(cookie)
}

func DeleteCookies(c *echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = "session_key"
	cookie.Value = ""
	cookie.Path = "/"
	cookie.HttpOnly = true
	cookie.Secure = true
	cookie.SameSite = http.SameSiteStrictMode
	cookie.MaxAge = -1
	cookie.Expires = time.Unix(1, 0)
	c.SetCookie(cookie)
}

// GetUserFromSession возвращает пользователя из сессии
// всё кроме ErrUnauthorized является ошибкой сервера,
// а ErrUnauthorized нужно обработать и выдать 401
// обычные эндпоинты должны использовать мидлварь OnlyUsersMiddleware.
func GetUserFromSession(c *echo.Context) (*database.User, error) {
	sessionStorage, err := GetSessionStorage(c)
	if err != nil {
		return nil, fmt.Errorf("failed to get user from session: %w", err)
	}

	sessionKey, err := GetKeyFromCookies(c)
	if err != nil {
		return nil, fmt.Errorf("failed to get user from session: %w", err)
	}

	userId, err := sessionStorage.GetIdFromSession(sessionKey)
	if err != nil {
		DeleteCookies(c)

		return nil, fmt.Errorf(
			"failed to get user from session: %w. You cookies is broken, we delete it",
			err,
		)
	}

	db, err := databaseMiddleware.GetDatabase(c)
	if err != nil {
		return nil, echo.NewHTTPError(
			http.StatusInternalServerError,
			fmt.Sprintf("failed to get database: %s.", err.Error()),
		)
	}

	user, err := db.GetUserByID(userId)
	if err != nil {
		DeleteCookies(c)

		return nil, echo.NewHTTPError(
			http.StatusInternalServerError,
			fmt.Sprintf(
				"can't get user by login: %s. You cookies is broken, we delete it",
				err.Error(),
			),
		)
	}

	return user, nil
}
