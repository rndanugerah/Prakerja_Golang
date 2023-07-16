package routes

import (
	"net/http"
	"prakerja/controllers/employees"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/employees", employees.FetchAllEmployees)
	e.GET("/employees/:id", employees.GetEmployeeByID)
	e.POST("/employees", employees.StoreEmployees)
	e.PUT("/employees/:id", employees.UpdateEmployees)
	e.DELETE("/employees/:id", employees.DeleteEmployees)

	return e
}
