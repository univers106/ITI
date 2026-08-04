package users

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/univers106/ITI/database"
	user_database_middleware "github.com/univers106/ITI/middlewares/databases/user"
)

type createUserRequest struct {
	Login       string   `json:"login" validate:"required,min=4"`
	Name        string   `json:"name" validate:"required,min=1"`
	Permissions []string `json:"permissions"`
	Password    string   `json:"password" validate:"required,min=4"`
}

func Post(c *echo.Context) error {
	user_db, err := user_database_middleware.Get(c)
	if err != nil {
		return fmt.Errorf("failed to get user database: %w", err)
	}

	var req createUserRequest
	err = c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	slog.Info("validating request", "req", req)
	err = c.Validate(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	newUser := database.User{
		Login:       req.Login,
		Name:        req.Name,
		Permissions: req.Permissions,
	}
	newUser.SetPassword(req.Password)

	err = user_db.Create(newUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "ok"})
}
