package main

import (
	"praktikum/config"
	"praktikum/route"
)

func main() {
	config.InitDB()
	e := route.New()
	e.Logger.Fatal(e.Start(":8000"))
}
