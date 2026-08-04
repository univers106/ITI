package user_database_middleware

import (
	"fmt"

	"github.com/labstack/echo/v5"
	"github.com/univers106/ITI/database"
)

func NewMiddleware(db database.UserDatabase) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			c.Set("_user_database", db)

			return next(c)
		}
	}
}

func Get(c *echo.Context) (database.UserDatabase, error) {
	database, err := echo.ContextGet[database.UserDatabase](c, "_user_database")
	if err != nil {
		return nil, fmt.Errorf("failed to get session store: %w", err)
	}

	return database, nil
}
