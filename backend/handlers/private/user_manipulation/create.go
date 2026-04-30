package user_manipulation

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/univers106/ITI/database"
	"github.com/univers106/ITI/middlewares/sessions_middleware"
)

type createUserRequest struct {
	Login    string `form:"login"`
	Name     string `form:"name"`
	Password string `form:"password"`
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

	err = db.CreateUser(request.Login, request.Name, request.Password)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			"failed to create user: "+err.Error(),
		)
	}

	return c.NoContent(http.StatusOK)
}
