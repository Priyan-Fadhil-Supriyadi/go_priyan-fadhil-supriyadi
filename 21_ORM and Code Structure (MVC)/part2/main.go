package main

import (
	"part2/config"
	"part2/route"
)

func main() {
	config.InitDB()
	e := route.New()
	e.Logger.Fatal(e.Start(":8000"))
}
