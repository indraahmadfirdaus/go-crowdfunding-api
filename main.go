package main

import (
	"crowdfunding-api/src/kernel"
	"crowdfunding-api/src/routes"
)

func main() {
	kernel.InitDB()
	routes := routes.NewRoutes()

	routes.Run()
}
