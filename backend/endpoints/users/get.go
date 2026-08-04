package users

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v5"
	user_database_middleware "github.com/univers106/ITI/middlewares/databases/user"
)

func Get(c *echo.Context) error {

	user_db, err := user_database_middleware.Get(c)
	if err != nil {
		return fmt.Errorf("failed to get user database: %w", err)
	}

	users, err := user_db.GetAll()
	if err != nil {
		return fmt.Errorf("failed to get all users: %w", err)
	}

	return c.JSON(http.StatusOK, users)
}
