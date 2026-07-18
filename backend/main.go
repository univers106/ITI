package main

import (
	"log/slog"
	"os"

	"github.com/labstack/echo/v5"
	echoMiddlewares "github.com/labstack/echo/v5/middleware"
	"github.com/univers106/ITI/config"
	"github.com/univers106/ITI/database"
	"github.com/univers106/ITI/database/postgresql"
	"github.com/univers106/ITI/database/postgresql/user_db"
	"github.com/univers106/ITI/handlers/private"
	"github.com/univers106/ITI/handlers/private/user_manipulation"
	"github.com/univers106/ITI/handlers/public"
	user_database_middleware "github.com/univers106/ITI/middlewares/database_middleware/user"
	"github.com/univers106/ITI/middlewares/sessions_middleware"
)

func main() {
	cfg := config.ReadConfig("config.yaml")

	// пока без конфига
	// postgres://[user]:[password]@[host]:[port]/[dbname]?[options]
	pgpool := postgresql.NewPool(cfg.PostgresSqlURL)

	var user_db database.UserDatabase = user_db.NewUserDatabase(pgpool)

	argsWithProg := os.Args
	if len(argsWithProg) > 1 {
		if argsWithProg[1] == "create-super-user" {
			user_db.CreateUser(
				database.User{
					Login:       argsWithProg[2],
					Name:        argsWithProg[2],
					Permissions: []string{database.PermSuperUser},
				},
				argsWithProg[3],
			)

		}
	}

	sessionStorage := sessions_middleware.NewSessionStorage()
	mainSessionMiddleware := sessions_middleware.NewSessionsMiddleware(sessionStorage)

	echoServer := echo.New()

	echoServer.Use(echoMiddlewares.RequestLogger())
	echoServer.Use(echoMiddlewares.Recover())
	echoServer.Use(user_database_middleware.NewMiddleware(user_db))

	apiGroup := echoServer.Group("/api")
	privateApi := apiGroup.Group("/private", mainSessionMiddleware)
	privateApi.Use(sessions_middleware.OnlyUsersMiddleware)

	privateApi.GET("/hello", private.GetHello)
	privateApi.GET("/logout", private.PostLogout)

	userManipulationApi := privateApi.Group("/user-manipulation")
	userManipulationApi.POST("/create", user_manipulation.PostCreate)
	userManipulationApi.POST("/delete", user_manipulation.PostDelete)
	userManipulationApi.POST("/change-password", user_manipulation.PostChangePassword)

	publicApi := apiGroup.Group("/public")

	publicApi.GET("/hello", public.GetHello)
	publicApi.POST("/login", public.PostLogin, mainSessionMiddleware)

	err := echoServer.Start(cfg.Host + ":8080")
	if err != nil {
		slog.Error("failed to start server", "error", err)
	}
}
