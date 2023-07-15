package employees

import (
	"net/http"
	"prakerja/models/employees"
	"github.com/labstack/echo/v4"
)

func FetchAllEmployees(c echo.Context) error {
	result, err := employees.FetchAllEmployees()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
