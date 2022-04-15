package controller

import "github.com/labstack/echo/v4"

func RegisterUserGroupAPI(e *echo.Group) {
	e.GET("", getUsersController)
	e.GET("/:id", getUserController)
	e.PUT("/:id", updateUserController)
	e.DELETE("/:id", deleteUserController)
	e.POST("", createUserController)
}