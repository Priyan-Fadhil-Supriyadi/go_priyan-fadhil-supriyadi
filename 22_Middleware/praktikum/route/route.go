package route

import (
	"praktikum/controller"
	m "praktikum/middleware"

	"github.com/labstack/echo/v4"
)

// create a new echo instance
func New() *echo.Echo {
	e := echo.New()
	// Route / to handler function
	m.LogMiddleware(e)
	e.GET("/users", controller.GetUsersController)
	e.GET("/users/:id", controller.GetUserController)
	e.POST("/users", controller.CreateUserController)
	e.DELETE("/users/:id", controller.DeleteUserController)
	e.PUT("/users/:id", controller.UpdateUserController)

	e.GET("/books", controller.GetBooksController)
	e.GET("/books/:id", controller.GetBookController)
	e.POST("/books", controller.CreateBookController)
	e.DELETE("/books/:id", controller.DeleteBookController)
	e.PUT("/books/:id", controller.UpdateBookController)

	return e
}
