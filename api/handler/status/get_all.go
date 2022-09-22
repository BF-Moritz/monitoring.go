package status

import (
	"net/http"
	"strconv"

	"github.com/BF-Moritz/monitoring.go/api/service/status"
	"github.com/labstack/echo/v4"
)

func NewGetAllHandler() echo.HandlerFunc {
	return newGetAllHandlerWithDeps(status.NewService())
}

func newGetAllHandlerWithDeps(statusService status.ServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		var lastMinutes uint32 = 1

		if c.QueryParam("last_minutes") != "" {
			parsedLastMinutes, err := strconv.ParseUint(c.QueryParam("last_minutes"), 10, 32)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, "'last_minutes' must be a number")
			}
			lastMinutes = uint32(parsedLastMinutes)
		}

		status, err := statusService.GetAll(lastMinutes)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, status)
	}
}
