package employees

import (
	"net/http"
	"prakerja/models/employees"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetEmployeeByID(c echo.Context) error {
	id := c.Param("id")

	idStr, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := employees.GetEmployeeByID(idStr)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
