package main

import (
	"restapi/config"
	"restapi/route"
)

func main() {
	config.InitDB()
	e := route.New()
	e.Logger.Fatal(e.Start(":8000"))
}
