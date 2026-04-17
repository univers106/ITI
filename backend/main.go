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
	"github.com/univers106/ITI/handlers/auth"
	"github.com/univers106/ITI/handlers/private"
	"github.com/univers106/ITI/handlers/public"
	databaseMiddleware "github.com/univers106/ITI/middlewares/database"
	sessionsMiddleware "github.com/univers106/ITI/middlewares/sessions"
)

func main() {
	cfg := config.ReadConfig("config.yaml")

	var db database.Database = filebased.NewFileBasedDatabase(cfg.DataDir)
	if _, err := db.GetUserByLogin("test_user"); errors.Is(err, database.ErrUserNotFound) {
		db.AddUser("test_user", "Test User", "test_password")
	}

	cookieStore := gorillaSessions.NewCookieStore([]byte(cfg.SessionKey))
	cookieStore.Options.Domain = cfg.Domain
	cookieStore.Options.Path = "/"
	cookieStore.Options.MaxAge = 60 * 60 * 24
	sessionMiddleware := sessionsMiddleware.NewSessionMiddleware(cookieStore)

	e := echo.New()

	e.Use(echoMiddlewares.RequestLogger())
	e.Use(echoMiddlewares.Recover())
	e.Use(databaseMiddleware.NewDatabaseMiddleware(db))

	apiGroup := e.Group("/api")
	privateApi := apiGroup.Group("/private", sessionMiddleware)
	publicApi := apiGroup.Group("/public")
	authApi := apiGroup.Group("/auth", sessionMiddleware)

	privateApi.GET("/hello", private.GetHello)
	publicApi.GET("/hello", public.GetHello)
	authApi.POST("/login", auth.PostLogin)

	if err := e.Start(":8080"); err != nil {
		slog.Error("failed to start server", "error", err)
	}
}
