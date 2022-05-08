package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"praktikum/config"
	"praktikum/database"
	"praktikum/domain"
	m "praktikum/middleware"
	"praktikum/model"
	"praktikum/repository"
)

func RegisterUserGroupAPI(e *echo.Echo, conf config.Config) {

	db := database.InitDB(conf)
	repo := repository.NewMysqlRepository(db)

	cont := NewEchoController(repo)

	apiUser := e.Group("/users",
		middleware.Logger(),
		middleware.CORS(),
	)
	apiUser.GET("", cont.GetUsersController)
	apiUser.GET("/:id", cont.GetUserController)
	apiUser.PUT("/:id", cont.UpdateUserController)
	apiUser.DELETE("/:id", cont.DeleteUserController)
	apiUser.POST("", cont.CreateUserController, middleware.BasicAuth(m.AuthUser))
}
