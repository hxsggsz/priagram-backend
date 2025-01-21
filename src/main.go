package main

import (
	"priagram/src/api/config"
	"priagram/src/api/routes"
)

// @host localhost:8080/api
// @BasePath /api
func main() {
	routes.InitializeRoutes()
	config.InitServer()
}
