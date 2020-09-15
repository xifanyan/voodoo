package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Init - initialize echo router
func Init(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

}
