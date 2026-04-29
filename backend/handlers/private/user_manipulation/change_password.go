package user_manipulation

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/univers106/ITI/database"
	"github.com/univers106/ITI/middlewares/sessions_middleware"
)

type changePasswordRequest struct {
	UserId   int    `form:"userId"`
	Password string `form:"password"`
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

	err = db.ChangeUserPassword(request.UserId, request.Password)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			"failed to change password: "+err.Error(),
		)
	}

	return c.NoContent(http.StatusOK)
}
