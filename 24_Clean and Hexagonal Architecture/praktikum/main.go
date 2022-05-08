package main

import (
	"github.com/labstack/echo/v4"

	conf "praktikum/config"
	rest "praktikum/controller"
)

func main() {
	config := conf.InitConfiguration()
	e := echo.New()

	rest.RegisterUserGroupAPI(e, config)

	e.Logger.Fatal(e.Start(":8080"))
}
