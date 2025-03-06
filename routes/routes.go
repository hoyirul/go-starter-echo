package routes

import (
	"go-echo/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Init initializes the routes
func Init() *echo.Echo {
	e := echo.New()

	// Root endpoint
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// User routes
	e.POST("/users", controllers.CreateUser)
	e.GET("/users", controllers.GetAllUsers)
	e.GET("/users/:id", controllers.GetUserByID)
	e.PUT("/users/:id", controllers.UpdateUser)
	e.DELETE("/users/:id", controllers.DeleteUser)

	return e
}
