package main

import (
	"log/slog"

	echojwt "github.com/labstack/echo-jwt/v5"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/univers106/ITI/config"
	"github.com/univers106/ITI/handlers/public"
)

func main() {
	cfg := config.ReadConfig("config.yaml")

	jwtMiddleware := echojwt.WithConfig(
		echojwt.Config{
			SigningKey: cfg.JwtKey,
		},
	)

	e := echo.New()

	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	apiGroup := e.Group("/api")
	privateApi := apiGroup.Group("/private", jwtMiddleware)
	publicApi := apiGroup.Group("/public")

	privateApi.GET("/hello", public.GetHello)
	publicApi.GET("/hello", public.GetHello)

	if err := e.Start(":8080"); err != nil {
		slog.Error("failed to start server", "error", err)
	}
}
