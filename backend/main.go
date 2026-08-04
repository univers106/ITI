package main

import (
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	echoMiddlewares "github.com/labstack/echo/v5/middleware"
	"github.com/univers106/ITI/cli"
	"github.com/univers106/ITI/config"
	"github.com/univers106/ITI/database"
	"github.com/univers106/ITI/database/postgresql"
	"github.com/univers106/ITI/database/postgresql/user_db"
	"github.com/univers106/ITI/endpoints/auth"
	"github.com/univers106/ITI/endpoints/users"
	user_database_middleware "github.com/univers106/ITI/middlewares/databases/user"
	sessions "github.com/univers106/ITI/middlewares/sessions"
	"github.com/univers106/ITI/middlewares/sessions/map_based"
)

type Validator struct {
	validator *validator.Validate
}

func (v *Validator) Validate(i any) error {
	err := v.validator.Struct(i)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return nil
}

func main() {
	cfg := config.ReadConfig("config.yaml")

	pgpool := postgresql.NewPool(cfg.PostgresURL)
	defer pgpool.Close()

	userDb := user_db.NewUserDatabase(pgpool)

	if cli.Run(pgpool) {
		return
	}

	sessionStorage := map_based.NewSessionStorage()
	mainSessionMiddleware := sessions.NewSessionsMiddleware(sessionStorage)

	echoServer := echo.New()

	echoServer.Use(echoMiddlewares.RequestLogger())
	echoServer.Use(echoMiddlewares.Recover())
	echoServer.Use(echoMiddlewares.RemoveTrailingSlash())
	echoServer.Use(user_database_middleware.NewMiddleware(userDb))

	apiGroup := echoServer.Group("/api")

	apiGroup.Use(mainSessionMiddleware)

	authGroup := apiGroup.Group("/auth")
	authGroup.POST("/login", auth.PostLogin)
	authGroup.POST("/logout", auth.PostLogout, sessions.Authed(nil))
	authGroup.PATCH("/change-password", auth.PatchChangePassword, sessions.Authed(nil))
	authGroup.GET("/me", auth.GetMe, sessions.Authed(nil))

	usersGroup := apiGroup.Group("/users")
	usersGroup.GET("", users.Get, sessions.Authed([]string{database.PermUsersManipulation}))

	echoServer.Validator = &Validator{validator: validator.New()}

	err := echoServer.Start(cfg.Host)
	if err != nil {
		slog.Error("failed to start server", "error", err)
	}
}
