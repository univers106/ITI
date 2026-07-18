package public

import (
	"net/http"

	"github.com/labstack/echo/v5"
	user_database_middleware "github.com/univers106/ITI/middlewares/database_middleware/user"
	"github.com/univers106/ITI/middlewares/sessions_middleware"
)

type loginRequest struct {
	Login    string `form:"userLogin" validate:"required,alphanum,min=2,max=32"`
	Password string `form:"password"  validate:"required,min=8,max=32,ascii"`
}

func PostLogin(c *echo.Context) error {
	sessionStorage, err := sessions_middleware.GetSessionStorage(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to get session")
	}

	_, err = sessions_middleware.GetKeyFromCookies(c)
	if err == nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			"session already exists, try logging out first",
		)
	}

	var req loginRequest

	err = c.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}

	err = c.Validate(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	db, err := user_database_middleware.Get(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to get database")
	}

	user, err := db.UserAuthentication(req.Login, req.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
	}

	sessionKey, err := sessionStorage.NewSession(user.Login)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to create session")
	}

	sessions_middleware.SetKeyToCookies(c, sessionKey)

	return c.JSON(http.StatusOK, map[string]string{"message": "ok"})
}
