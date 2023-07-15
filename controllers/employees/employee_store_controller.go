package employees

import (
	"net/http"
	"prakerja/models/employees"

	"github.com/labstack/echo/v4"
)

func StoreEmployees(c echo.Context) error {
	nama := c.FormValue("nama")
	email := c.FormValue("email")
	telepon := c.FormValue("telepon")

	result, err := employees.StoreEmployees(nama, email, telepon)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
