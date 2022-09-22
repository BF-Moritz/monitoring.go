package router

import (
	"github.com/BF-Moritz/monitoring.go/api/handler/status"
	"github.com/labstack/echo/v4"
)

func makeStatusRoutes(g *echo.Group) {
	g.GET("/", status.NewGetAllHandler())
	g.GET("/:name", status.NewGetByNameHandler())
}
