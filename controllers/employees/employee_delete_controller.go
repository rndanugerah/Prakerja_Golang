package employees

import (
	"net/http"
	"prakerja/models/employees"
	"strconv"

	"github.com/labstack/echo/v4"
)

func DeleteEmployees(c echo.Context) error {
	id := c.Param("id")

	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Empty 'id' parameter. It should be a valid integer.",
		})
	}

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := employees.DeleteEmployees(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
