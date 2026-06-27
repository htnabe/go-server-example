package main

import (
	"github.com/htnabe/go-server-example/backend/internal/router"
	"github.com/labstack/echo/v5/middleware"
)

func main() {
	e := router.New()
	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	if err := e.Start(":1323"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
