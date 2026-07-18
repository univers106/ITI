package user_manipulation

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/univers106/ITI/database"
	"github.com/univers106/ITI/middlewares/sessions_middleware"
)

type changeNameRequest struct {
	UserLogin string `form:"userLogin" validate:"required,alphanum,min=2,max=32"`
	Name      string `form:"name"      validate:"required,min=2,max=50"`
}

func PostChangeName(c *echo.Context) error {
	_, db, httpErr := sessions_middleware.GetUserDbCheckPermision(
		c,
		database.PermUsersManipulation,
	)
	if httpErr != nil {
		return httpErr
	}

	var request changeNameRequest

	err := c.Bind(&request)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "There is something wrong with the values")
	}

	err = c.Validate(&request)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = db.ChangeUserName(request.UserLogin, request.Name)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			"failed to change name",
		)
	}

	return c.NoContent(http.StatusOK)
}
