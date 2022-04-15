package main

import (
	"github.com/labstack/echo/v4"

	conf "praktikum/config"
	c "praktikum/controller"
	db "praktikum/database"
)

func init() {
	config := conf.InitConfiguration()
	db.InitDB(config)
}

func main() {
	e := echo.New()
	apiUser := e.Group("/users")
	c.RegisterUserGroupAPI(apiUser)

	e.Logger.Fatal(e.Start(":8080"))
}
