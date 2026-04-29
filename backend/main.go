package main

import (
	"errors"
	"log/slog"

	"github.com/labstack/echo/v5"
	echoMiddlewares "github.com/labstack/echo/v5/middleware"
	"github.com/univers106/ITI/config"
	"github.com/univers106/ITI/database"
	"github.com/univers106/ITI/database/file_based"
	"github.com/univers106/ITI/handlers/private"
	"github.com/univers106/ITI/handlers/private/user_manipulation"
	"github.com/univers106/ITI/handlers/public"
	"github.com/univers106/ITI/middlewares/database_middleware"
	"github.com/univers106/ITI/middlewares/sessions_middleware"
)

func main() {
	cfg := config.ReadConfig("config.yaml")

	var db database.Database = file_based.NewFileBasedDatabase(cfg.DataDir)

	// временно

	_, err := db.GetUserByLogin("test_user")
	if errors.Is(err, database.ErrUserNotFound) {
		//nolint
		db.CreateUser("test_user", "Test User", "test_password")
	}

	_, err = db.GetUserByLogin("test_admin")
	if errors.Is(err, database.ErrUserNotFound) {
		//nolint
		db.CreateUser("test_admin", "ADMIN", "test_password")

		admin, err := db.GetUserByLogin("test_admin")
		if err != nil {
			panic(err)
		}

		//nolint
		db.UserAddPermissions(admin.ID, database.PermSuperUser)
	}

	// конец временно

	sessionStorage := sessions_middleware.NewSessionStorage()
	mainSessionMiddleware := sessions_middleware.NewSessionsMiddleware(sessionStorage)

	echoServer := echo.New()

	echoServer.Use(echoMiddlewares.RequestLogger())
	echoServer.Use(echoMiddlewares.Recover())
	echoServer.Use(database_middleware.NewDatabaseMiddleware(db))

	apiGroup := echoServer.Group("/api")
	privateApi := apiGroup.Group("/private", mainSessionMiddleware)
	privateApi.Use(sessions_middleware.OnlyUsersMiddleware)

	publicApi := apiGroup.Group("/public")

	privateApi.GET("/hello", private.GetHello)
	privateApi.GET("/logout", private.PostLogout)

	userManipulationApi := privateApi.Group("/user-manipulation")
	userManipulationApi.POST("/create", user_manipulation.PostCreate)
	userManipulationApi.POST("/delete", user_manipulation.PostDelete)
	userManipulationApi.POST("/change-password", user_manipulation.PostChangePassword)

	publicApi.GET("/hello", public.GetHello)
	publicApi.POST("/login", public.PostLogin, mainSessionMiddleware)

	err = echoServer.Start(":8080")
	if err != nil {
		slog.Error("failed to start server", "error", err)
	}
}
