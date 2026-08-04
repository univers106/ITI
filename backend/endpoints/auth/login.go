package auth

import (
	"net/http"

	"github.com/labstack/echo/v5"
	user_database_middleware "github.com/univers106/ITI/middlewares/databases/user"
	"github.com/univers106/ITI/middlewares/sessions"
)

type loginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func PostLogin(c *echo.Context) error {
	var req loginRequest

	err := c.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}

	if req.Password == "" || req.Login == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "login or password value is null")
	}

	db, err := user_database_middleware.Get(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to get database")
	}

	user, err := db.UserAuthentication(req.Login, req.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
	}

	err = sessions.NewSession(c, user.Login)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to create session")
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "ok"})
}
