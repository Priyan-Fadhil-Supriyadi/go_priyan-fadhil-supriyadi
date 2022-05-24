package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/Priyan-Fadhil-Supriyadi/mini-project/config"
	"github.com/Priyan-Fadhil-Supriyadi/mini-project/database"
	m "github.com/Priyan-Fadhil-Supriyadi/mini-project/middleware"
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

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"message": "your request awesome",
		})
	})

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
	apiUser.POST("", cont.CreateUserController)
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
	apiCar.PUT("/:id", cont.UpdateCarController, middleware.BasicAuth(m.AuthUser))
	apiCar.DELETE("/:id", cont.DeleteCarController, middleware.BasicAuth(m.AuthUser))
	apiCar.POST("", cont.CreateCarController, middleware.BasicAuth(m.AuthUser))
}

func DepotGroupAPI(e *echo.Echo, conf config.Config) {

	db := database.InitDB(conf)

	repoDepot := repository.DepotMysqlRepository(db)

	svcDepot := service.NewServiceDepot(repoDepot, conf)

	cont := DepotController{
		svc: svcDepot,
	}

	apiDepot := e.Group("/depots",
		middleware.Logger(),
		middleware.CORS(),
		//m.APIKEYMiddleware,
	)

	apiDepot.GET("", cont.GetDepotsController, middleware.JWTWithConfig(
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

	// apiDepot.GET("", cont.GetDepotsController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiDepot.GET("/:id", cont.GetDepotController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiDepot.PUT("/:id", cont.UpdateDepotController, middleware.BasicAuth(m.AuthUser))
	apiDepot.DELETE("/:id", cont.DeleteDepotController, middleware.BasicAuth(m.AuthUser))
	apiDepot.POST("", cont.CreateDepotController, middleware.BasicAuth(m.AuthUser))
}

func RentGroupAPI(e *echo.Echo, conf config.Config) {

	db := database.InitDB(conf)

	repoRent := repository.RentMysqlRepository(db)

	svcRent := service.NewServiceRent(repoRent, conf)

	cont := RentController{
		svc: svcRent,
	}

	apiRent := e.Group("/rents",
		middleware.Logger(),
		middleware.CORS(),
		//m.APIKEYMiddleware,
	)

	apiRent.GET("", cont.GetRentsController, middleware.JWTWithConfig(
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

	// apiRent.GET("", cont.GetRentsController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiRent.GET("/:id", cont.GetRentController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiRent.PUT("/:id", cont.UpdateRentController, middleware.JWTWithConfig(
		middleware.JWTConfig{
			SigningKey: []byte(conf.JWT_KEY),
		},
	))
	apiRent.DELETE("/:id", cont.DeleteRentController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiRent.POST("", cont.CreateRentController, middleware.JWT([]byte(conf.JWT_KEY)))
}
