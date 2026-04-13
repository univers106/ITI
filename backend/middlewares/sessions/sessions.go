package sessionsMiddleware

import (
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
