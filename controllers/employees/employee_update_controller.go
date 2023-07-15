package employees

import (
	"net/http"
	"prakerja/models/employees"
	"strconv"

	"github.com/labstack/echo/v4"
)

func UpdateEmployees(c echo.Context) error {
	id := c.FormValue("id")
	nama := c.FormValue("nama")
	email := c.FormValue("email")
	telepon := c.FormValue("telepon")

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := employees.UpdateEmployees(conv_id, nama, email, telepon)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
