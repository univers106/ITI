package main

import (
	"errors"
	"log/slog"
	"time"

	gorillaSessions "github.com/gorilla/sessions"
	"github.com/labstack/echo/v5"
	echoMiddlewares "github.com/labstack/echo/v5/middleware"
	"github.com/univers106/ITI/config"
	"github.com/univers106/ITI/database"
	filebased "github.com/univers106/ITI/database/file_based"
	"github.com/univers106/ITI/handlers/private"
	"github.com/univers106/ITI/handlers/public"
	"github.com/univers106/ITI/middlewares/databaseMiddleware"
	"github.com/univers106/ITI/middlewares/sessionsMiddleware"
)

func main() {
	cfg := config.ReadConfig("config.yaml")

	var db database.Database = filebased.NewFileBasedDatabase(cfg.DataDir)

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

	cookieStore := gorillaSessions.NewCookieStore([]byte(cfg.SessionKey))
	cookieStore.Options.Domain = cfg.Domain
	cookieStore.Options.Path = "/"
	cookieStore.Options.MaxAge = int(time.Hour.Seconds())
	mainSessionMiddleware := sessionsMiddleware.NewSessionsMiddleware(cookieStore)

	echoServer := echo.New()

	echoServer.Use(echoMiddlewares.RequestLogger())
	echoServer.Use(echoMiddlewares.Recover())
	echoServer.Use(databaseMiddleware.NewDatabaseMiddleware(db))

	apiGroup := echoServer.Group("/api")
	privateApi := apiGroup.Group("/private", mainSessionMiddleware)
	privateApi.Use(sessionsMiddleware.OnlyUsersMiddleware)

	publicApi := apiGroup.Group("/public")

	privateApi.GET("/hello", private.GetHello)
	privateApi.GET("/logout", private.PostLogout)
	privateApi.POST("/create-user", private.PostCreateUser)

	publicApi.GET("/hello", public.GetHello)
	publicApi.POST("/login", public.PostLogin, mainSessionMiddleware)

	err = echoServer.Start(":8080")
	if err != nil {
		slog.Error("failed to start server", "error", err)
	}
}
