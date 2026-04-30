package user_manipulation

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/univers106/ITI/database"
	"github.com/univers106/ITI/middlewares/sessions_middleware"
)

type changeLoginRequest struct {
	userId int    `form:"userId"`
	login  string `form:"login"`
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

	err = db.ChangeUserLogin(request.userId, request.login)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			"failed to change login: "+err.Error(),
		)
	}

	return c.NoContent(http.StatusOK)
}
