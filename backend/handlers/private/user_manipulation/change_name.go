package user_manipulation

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/univers106/ITI/database"
	"github.com/univers106/ITI/middlewares/sessions_middleware"
)

type changeNameRequest struct {
	UserId int    `form:"userId"`
	Name   string `form:"name"`
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

	err = db.ChangeUserName(request.UserId, request.Name)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			"failed to change name: "+err.Error(),
		)
	}

	return c.NoContent(http.StatusOK)
}
