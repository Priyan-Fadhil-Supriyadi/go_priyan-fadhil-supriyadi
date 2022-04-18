package main

import (
	"github.com/labstack/echo/v4"

	conf "github.com/bimbimprasetyoafif/km/config"
	c "github.com/bimbimprasetyoafif/km/controller"
	db "github.com/bimbimprasetyoafif/km/database"
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












/*


- cmd/main.go
- config = konfigurasi apapun -> membaca env
- db
- route = rest endpoint / endpoint doang
- controller = routing (sedikit logic kalau perlu)
- model = model (database), model struct
- repository = logic database (rdbms / query)
- lib / helper / util = pembantu



- cmd/main.go // go run cmd/main.go
- internal / pkg
	- route = rest endpoint / endpoint doang
	- controller = routing (sedikit logic kalau perlu)
	- repository = logic database (rdbms / query)
- model = model (database), model struct
- lib / helper / util = pembantu



- cmd/cmd.go // go run cmd/main.go
- config
- internal
	- api
		- rest
			- route = rest endpoint / endpoint doang
			- controller = routing (sedikit logic kalau perlu)
		- grpc
			- controller = routing (sedikit logic kalau perlu)
		- graphql
			- route = rest endpoint / endpoint doang
			- controller = routing (sedikit logic kalau perlu)

	- server
		-rest
			-rest.go
		-grpc
			-grpc.go
		-graphql
			-graphql.go

	- repository = logic database (rdbms / query)

main.go








- vendor = pihak 3 / nodmodules
- app/Controller/cafauhfa.php
- app/Model
- config
- route
- public/
	- view
	- *.php




















*/