package main

import (
	"VPNService/app"
	"fmt"
)

func main() {

	app := app.New()

	fmt.Println("App was creating")

	app.Run()

	//init repository
	//init services
	//init api

	//start server

	//server loop

	//graceful shutdown
}
