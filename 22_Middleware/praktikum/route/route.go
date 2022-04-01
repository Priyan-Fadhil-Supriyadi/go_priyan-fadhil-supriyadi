package route

import (
	"praktikum/constants"
	"praktikum/controller"
	m "praktikum/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// create a new echo instance
func New() *echo.Echo {
	e := echo.New()

	// middleware
	m.LogMiddleware(e)
	// eAuth := e.Group("")
	// eAuth.Use(middleware.BasicAuth(m.BasicAuthDB))

	eJwt := e.Group("")
	eJwt.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	// Route / to handler function
	e.POST("/login", controller.LoginUserController)

	eJwt.GET("/users", controller.GetUsersController)
	eJwt.GET("/users/:id", controller.GetUserController)
	e.POST("/users", controller.CreateUserController)
	eJwt.DELETE("/users/:id", controller.DeleteUserController)
	eJwt.PUT("/users/:id", controller.UpdateUserController)

	e.GET("/books", controller.GetBooksController)
	e.GET("/books/:id", controller.GetBookController)
	eJwt.POST("/books", controller.CreateBookController)
	eJwt.DELETE("/books/:id", controller.DeleteBookController)
	eJwt.PUT("/books/:id", controller.UpdateBookController)

	return e
}
