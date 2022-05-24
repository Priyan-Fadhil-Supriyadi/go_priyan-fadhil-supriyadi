package main

import (
	"github.com/labstack/echo/v4"

	conf "github.com/Priyan-Fadhil-Supriyadi/mini-project/config"
	rest "github.com/Priyan-Fadhil-Supriyadi/mini-project/controller"
)

func main() {
	config := conf.InitConfiguration()
	e := echo.New()

	rest.UserGroupAPI(e, config)
	rest.CarGroupAPI(e, config)
	rest.DepotGroupAPI(e, config)
	rest.RentGroupAPI(e, config)

	e.Logger.Fatal(e.Start(config.SERVER_ADDRESS))
}
