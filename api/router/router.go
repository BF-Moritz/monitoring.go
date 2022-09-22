package router

import (
	"github.com/labstack/echo/v4"
)

func MakeRoutes(e *echo.Echo) {
	g := e.Group("/api/v1")

	makeStatusRoutes(g.Group("/status"))
}
