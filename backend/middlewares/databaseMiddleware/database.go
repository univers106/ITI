package databaseMiddleware

import (
	"fmt"

	"github.com/labstack/echo/v5"
	"github.com/univers106/ITI/database"
)

func NewDatabaseMiddleware(db database.Database) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			c.Set("_database", db)

			return next(c)
		}
	}
}

func GetDatabase(c *echo.Context) (database.Database, error) {
	database, err := echo.ContextGet[database.Database](c, "_database")
	if err != nil {
		return nil, fmt.Errorf("failed to get session store: %w", err)
	}

	return database, nil
}
