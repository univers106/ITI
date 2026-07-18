package main

import (
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	echoMiddlewares "github.com/labstack/echo/v5/middleware"
	"github.com/univers106/ITI/cli"
	"github.com/univers106/ITI/config"
	"github.com/univers106/ITI/database/postgresql"
	"github.com/univers106/ITI/database/postgresql/user_db"
	"github.com/univers106/ITI/handlers/private"
	"github.com/univers106/ITI/handlers/private/user_manipulation"
	"github.com/univers106/ITI/handlers/public"
	user_database_middleware "github.com/univers106/ITI/middlewares/database_middleware/user"
	"github.com/univers106/ITI/middlewares/sessions_middleware"
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

	// пока без конфига
	// postgres://[user]:[password]@[host]:[port]/[dbname]?[options]
	pgpool := postgresql.NewPool(cfg.PostgresURL)
	defer pgpool.Close()

	userDb := user_db.NewUserDatabase(pgpool)

	if cli.Run(pgpool) {
		return
	}

	sessionStorage := sessions_middleware.NewSessionStorage()
	mainSessionMiddleware := sessions_middleware.NewSessionsMiddleware(sessionStorage)

	echoServer := echo.New()

	echoServer.Use(echoMiddlewares.RequestLogger())
	echoServer.Use(echoMiddlewares.Recover())
	echoServer.Use(user_database_middleware.NewMiddleware(userDb))

	apiGroup := echoServer.Group("/api")
	privateApi := apiGroup.Group("/private", mainSessionMiddleware)
	privateApi.Use(sessions_middleware.OnlyUsersMiddleware)

	privateApi.GET("/hello", private.GetHello)
	privateApi.GET("/logout", private.GetLogout)

	userManipulationApi := privateApi.Group("/user-manipulation")
	userManipulationApi.POST("/create", user_manipulation.PostCreate)
	userManipulationApi.POST("/delete", user_manipulation.PostDelete)
	userManipulationApi.POST("/change-password", user_manipulation.PostChangePassword)

	publicApi := apiGroup.Group("/public")

	publicApi.GET("/hello", public.GetHello)
	publicApi.POST("/login", public.PostLogin, mainSessionMiddleware)

	echoServer.Validator = &Validator{validator: validator.New()}

	err := echoServer.Start(cfg.Host)
	if err != nil {
		slog.Error("failed to start server", "error", err)
	}
}
