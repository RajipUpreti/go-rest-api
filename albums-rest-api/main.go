package main

import (
	"albums-rest-api/config"
	"albums-rest-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.ConnectDatabase()

	routes.AlbumRoutes(r)
	routes.HealthCheckRoutes(r)
	r.Run()
}
