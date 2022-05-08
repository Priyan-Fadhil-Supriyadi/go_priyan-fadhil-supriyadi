package main

import (
	"github.com/labstack/echo/v4"

	conf "github.com/Priyan-Fadhil-Supriyadi/mini-project/config"
	rest "github.com/Priyan-Fadhil-Supriyadi/mini-project/controller"
)

func main() {
	config := conf.InitConfiguration()
	e := echo.New()

	rest.RegisterUserGroupAPI(e, config)

	e.Logger.Fatal(e.Start(":8080"))
}
