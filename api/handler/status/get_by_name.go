package status

import "github.com/labstack/echo/v4"

func NewGetByNameHandler() echo.HandlerFunc {
	return newGetByNameHandlerWithDeps()
}

func newGetByNameHandlerWithDeps() echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}
