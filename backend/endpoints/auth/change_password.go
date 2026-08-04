package auth

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v5"
	user_database_middleware "github.com/univers106/ITI/middlewares/databases/user"
	"github.com/univers106/ITI/middlewares/sessions"
)

type changePasswordRequest struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

func PatchChangePassword(c *echo.Context) error {
	var req changePasswordRequest

	user, err := sessions.GetUser(c)
	if err != nil {
		return fmt.Errorf("failed to get user from context: %w", err)
	}

	err = c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Wrong request body"})
	}

	user_db, err := user_database_middleware.Get(c)
	if err != nil {
		return fmt.Errorf("failed to get user from database: %w", err)
	}

	_, err = user_db.UserAuthentication(user.Login, req.OldPassword)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Wrong old password"})
	}

	user.SetPassword(req.NewPassword)
	err = user_db.Change(user.Login, *user)
	if err != nil {
		return fmt.Errorf("failed to change password: %w", err)
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "ok"})
}
