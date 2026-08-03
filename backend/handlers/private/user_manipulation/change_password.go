package user_manipulation

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/univers106/ITI/database"
	"github.com/univers106/ITI/middlewares/sessions_middleware"
)

type changePasswordRequest struct {
	UserLogin string `form:"userLogin" validate:"required,alphanum,min=2,max=32"`
	Password  string `form:"password"  validate:"required,min=8,max=32,ascii"`
}

func PostChangePassword(c *echo.Context) error {
	_, db, httpErr := sessions_middleware.GetUserDbCheckPermision(
		c,
		database.PermUsersManipulation,
	)
	if httpErr != nil {
		return httpErr
	}

	var request changePasswordRequest

	err := c.Bind(&request)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "There is something wrong with the values")
	}

	err = c.Validate(&request)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = db.ChangeUserPassword(request.UserLogin, request.Password)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			"failed to change password",
		)
	}

	return c.NoContent(http.StatusOK)
}
