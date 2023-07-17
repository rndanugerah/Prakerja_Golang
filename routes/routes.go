package routes

import (
	"net/http"
	"prakerja/controllers/employees"
	"prakerja/controllers/login"
	"prakerja/middleware"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/employees", employees.FetchAllEmployees)
	e.GET("/employees/:id", employees.GetEmployeeByID)
	e.POST("/employees", employees.StoreEmployees, middleware.IsAuthenticated)
	e.PUT("/employees/:id", employees.UpdateEmployees, middleware.IsAuthenticated)
	e.DELETE("/employees/:id", employees.DeleteEmployees, middleware.IsAuthenticated)

	e.GET("/generate-hash", login.GenerateHashPassword)
	e.POST("/login", login.CheckLogin)

	return e
}
