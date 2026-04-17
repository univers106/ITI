package sessionsMiddleware

import (
	"errors"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v5"
	"github.com/univers106/ITI/database"
	databaseMiddleware "github.com/univers106/ITI/middlewares/database"
)

const AuthSession = "auth"

var (
	ErrUnauthorized = errors.New("unauthorized")
	ErrBadCookies   = errors.New("bad cookies")
	ErrUserNotFound = errors.New("user not found")
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
		return nil, err
	}
	return store.Get(c.Request(), AuthSession)
}

// всё кроме ErrUnauthorized является ошибкой сервера,
// а ErrUnauthorized нужно обработать и выдать 401
// обычные эндпоинты должны использовать мидлварь OnlyUsersMiddleware
func GetUserFromSessionContext(c *echo.Context) (database.User, error) {
	session, err := GetAuthSessionFromContext(c)
	if err != nil {
		return database.User{}, err
	}

	login := session.Values["login"]
	if login == nil {
		return database.User{}, ErrUnauthorized
	}
	loginStr := login.(string)
	db, err := databaseMiddleware.GetDatabase(c)
	if err != nil {
		return database.User{}, err
	}

	user, err := db.GetUserByLogin(loginStr)
	if err != nil {
		return database.User{}, ErrUserNotFound
	}
	return user, nil
}
