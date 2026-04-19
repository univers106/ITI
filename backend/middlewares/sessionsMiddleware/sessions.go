package sessionsMiddleware

import (
	"errors"
	"fmt"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v5"
	"github.com/univers106/ITI/database"
	"github.com/univers106/ITI/middlewares/databaseMiddleware"
)

const AuthSession = "auth"

var (
	ErrUnauthorized = errors.New("unauthorized")
	ErrBadCookies   = errors.New("bad cookies")
)

func NewSessionsMiddleware(store sessions.Store) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			c.Set("_session_store", store)

			return next(c)
		}
	}
}

func GetAuthSessionFromContext(c *echo.Context) (*sessions.Session, error) {
	store, err := echo.ContextGet[sessions.Store](c, "_session_store")
	if err != nil {
		return nil, fmt.Errorf("failed to get session store: %w", err)
	}

	return store.Get(c.Request(), AuthSession)
}

// GetUserFromSessionContext возвращает пользователя из сессии
// всё кроме ErrUnauthorized является ошибкой сервера,
// а ErrUnauthorized нужно обработать и выдать 401
// обычные эндпоинты должны использовать мидлварь OnlyUsersMiddleware.
func GetUserFromSessionContext(c *echo.Context) (*database.User, error) {
	session, err := GetAuthSessionFromContext(c)
	if err != nil {
		return nil, err
	}

	login := session.Values["login"]
	if login == nil {
		return nil, ErrUnauthorized
	}

	loginStr, ok := login.(string)
	if !ok {
		return nil, ErrBadCookies
	}

	db, err := databaseMiddleware.GetDatabase(c)
	if err != nil {
		return nil, fmt.Errorf("failed to get database: %w", err)
	}

	user, err := db.GetUserByLogin(loginStr)
	if err != nil {
		return nil, fmt.Errorf("can't get user by login: %w", err)
	}

	return user, nil
}
