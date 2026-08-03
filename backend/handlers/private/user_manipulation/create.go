package user_manipulation

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/univers106/ITI/database"
	"github.com/univers106/ITI/middlewares/sessions_middleware"
)

type createUserRequest struct {
	Login    string `form:"login"    validate:"required,alphanum,min=2,max=32"`
	Name     string `form:"name"     validate:"required,min=2,max=50"`
	Password string `form:"password" validate:"required,min=8,max=32,ascii"`
}

func PostCreate(c *echo.Context) error {
	_, db, httpErr := sessions_middleware.GetUserDbCheckPermision(
		c,
		database.PermUsersManipulation,
	)
	if httpErr != nil {
		return httpErr
	}

	var request createUserRequest

	err := c.Bind(&request)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "There is something wrong with the values")
	}

	err = c.Validate(&request)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "There is something wrong with the values")
	}

	err = db.CreateUser(
		database.User{
			Login: request.Login,
			Name:  request.Name,
		},
		request.Password,
	)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			"failed to create user",
		)
	}

	return c.NoContent(http.StatusOK)
}
