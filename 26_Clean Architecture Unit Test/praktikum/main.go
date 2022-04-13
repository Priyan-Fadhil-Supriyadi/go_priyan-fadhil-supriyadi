package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	conf "praktikum/config"
	c "praktikum/controller"
	db "praktikum/database"
	m "praktikum/middleware"
)

func init() {
	config := conf.InitConfiguration()
	db.InitDB(config)
}

func main() {
	e := echo.New()
	apiUser := e.Group("/users",
		middleware.Logger(),
		middleware.CORS(),
		m.APIKEYMiddleware,
	)

	c.RegisterUserGroupAPI(apiUser)

	e.Logger.Fatal(e.Start(":8080"))
}
