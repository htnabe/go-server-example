package router

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Hello from Router!",
		})
	})

	return e
}
