package main

import (
	"errors"
	"log/slog"

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
	if _, err := db.GetUserByLogin("test_user"); errors.Is(err, database.ErrUserNotFound) {
		db.CreateUser("test_user", "Test User", "test_password")
	}
	if _, err := db.GetUserByLogin("test_admin"); errors.Is(err, database.ErrUserNotFound) {
		db.CreateUser("test_admin", "ADMIN", "test_password")
		admin, err := db.GetUserByLogin("test_admin")
		if err != nil {
			panic(err)
		}
		db.UserAddPermissions(admin.ID, database.PermSuperUser)
	}

	cookieStore := gorillaSessions.NewCookieStore([]byte(cfg.SessionKey))
	cookieStore.Options.Domain = cfg.Domain
	cookieStore.Options.Path = "/"
	cookieStore.Options.MaxAge = 60 * 60 * 1
	mainSessionMiddleware := sessionsMiddleware.NewSessionsMiddleware(cookieStore)

	e := echo.New()

	e.Use(echoMiddlewares.RequestLogger())
	e.Use(echoMiddlewares.Recover())
	e.Use(databaseMiddleware.NewDatabaseMiddleware(db))

	apiGroup := e.Group("/api")
	privateApi := apiGroup.Group("/private", mainSessionMiddleware)
	privateApi.Use(sessionsMiddleware.OnlyUsersMiddleware)
	publicApi := apiGroup.Group("/public")

	privateApi.GET("/hello", private.GetHello)
	privateApi.GET("/logout", private.PostLogut, mainSessionMiddleware)

	publicApi.GET("/hello", public.GetHello)
	publicApi.POST("/login", public.PostLogin, mainSessionMiddleware)

	if err := e.Start(":8080"); err != nil {
		slog.Error("failed to start server", "error", err)
	}
}
