package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/Priyan-Fadhil-Supriyadi/mini-project/config"
	"github.com/Priyan-Fadhil-Supriyadi/mini-project/database"

	//m "github.com/Priyan-Fadhil-Supriyadi/mini-project/middleware"

	"github.com/Priyan-Fadhil-Supriyadi/mini-project/repository"
	"github.com/Priyan-Fadhil-Supriyadi/mini-project/service"
)

func UserGroupAPI(e *echo.Echo, conf config.Config) {

	db := database.InitDB(conf)

	repo := repository.UserMysqlRepository(db)

	svcUser := service.NewServiceUser(repo, conf)

	cont := EchoController{
		svc: svcUser,
	}

	apiUser := e.Group("/users",
		middleware.Logger(),
		middleware.CORS(),
		//m.APIKEYMiddleware,
	)

	apiUser.GET("", cont.GetUsersController, middleware.JWTWithConfig(
		middleware.JWTConfig{
			SigningKey: []byte(conf.JWT_KEY),
			ErrorHandlerWithContext: func(err error, c echo.Context) error {
				return c.JSONPretty(404, map[string]interface{}{
					"messages": "token mu mungkin invalid",
				}, "  ")
			},
			SuccessHandler: func(c echo.Context) {
			},
		},
	))

	// apiUser.GET("", cont.GetUsersController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiUser.POST("/login", cont.LoginUserController)
	apiUser.GET("/:id", cont.GetUserController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiUser.PUT("/:id", cont.UpdateUserController, middleware.JWTWithConfig(
		middleware.JWTConfig{
			SigningKey: []byte(conf.JWT_KEY),
		},
	))
	apiUser.DELETE("/:id", cont.DeleteUserController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiUser.POST("", cont.CreateUserController, middleware.JWT([]byte(conf.JWT_KEY)))
}

func CarGroupAPI(e *echo.Echo, conf config.Config) {

	db := database.InitDB(conf)

	repoCar := repository.CarMysqlRepository(db)

	svcCar := service.NewServiceCar(repoCar, conf)

	cont := CarController{
		svc: svcCar,
	}

	apiCar := e.Group("/cars",
		middleware.Logger(),
		middleware.CORS(),
		//m.APIKEYMiddleware,
	)

	apiCar.GET("", cont.GetCarsController, middleware.JWTWithConfig(
		middleware.JWTConfig{
			SigningKey: []byte(conf.JWT_KEY),
			ErrorHandlerWithContext: func(err error, c echo.Context) error {
				return c.JSONPretty(404, map[string]interface{}{
					"messages": "token mu mungkin invalid",
				}, "  ")
			},
			SuccessHandler: func(c echo.Context) {
			},
		},
	))

	// apiCar.GET("", cont.GetCarsController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiCar.GET("/:id", cont.GetCarController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiCar.PUT("/:id", cont.UpdateCarController, middleware.JWTWithConfig(
		middleware.JWTConfig{
			SigningKey: []byte(conf.JWT_KEY),
		},
	))
	apiCar.DELETE("/:id", cont.DeleteCarController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiCar.POST("", cont.CreateCarController, middleware.JWT([]byte(conf.JWT_KEY)))
}
