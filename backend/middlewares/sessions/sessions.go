package sessionsMiddleware

import (
	"errors"
	"fmt"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v5"
)

const AuthSession = "auth"

func NewSessionMiddleware(store sessions.Store) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			c.Set("_session_store", store)
			return next(c)
		}
	}
}

func GetUserSession(c *echo.Context) (*sessions.Session, error) {
	store, err := echo.ContextGet[sessions.Store](c, "_session_store")
	if errors.Is(err, securecookie.ErrMacInvalid.Cause()) {
		c.JSON(403, map[string]string{"message": "bad cookies"})
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get session store: %w", err)
	}
	return store.Get(c.Request(), AuthSession)
}
