package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// RunApplication()
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))
	e.GET("/", func(ctx echo.Context) error {
		return ctx.JSON(200, map[string]interface{}{"msg": "hello"})
	})
	e.GET("/hello", func(ctx echo.Context) error {
		return ctx.JSON(200, map[string]interface{}{"msg": "hello"})
	})
	e.Start(":8080")
}
