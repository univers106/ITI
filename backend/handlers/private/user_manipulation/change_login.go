package user_manipulation

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/univers106/ITI/database"
	"github.com/univers106/ITI/middlewares/sessions_middleware"
)

type changeLoginRequest struct {
	UserLogin string `form:"userLogin" validate:"required,alphanum,min=2,max=32"`
	Login     string `form:"login"     validate:"required,alphanum,min=2,max=32"`
}

func PostChangeLogin(c *echo.Context) error {
	_, db, httpErr := sessions_middleware.GetUserDbCheckPermision(
		c,
		database.PermUsersManipulation,
	)
	if httpErr != nil {
		return httpErr
	}

	var request changeLoginRequest

	err := c.Bind(&request)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "There is something wrong with the values")
	}

	err = c.Validate(&request)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "There is something wrong with the values")
	}

	err = db.ChangeUserLogin(request.UserLogin, request.Login)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			"failed to change login",
		)
	}

	return c.NoContent(http.StatusOK)
}
