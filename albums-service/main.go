package main

import (
	"albums-service/config"
	"albums-service/models"
	"albums-service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.ConnectDatabase()
	if err := models.Migrate(config.DB); err != nil {
		panic(err)
	}
	routes.AlbumRoutes(r)
	routes.HealthCheckRoutes(r)
	r.Run()
}
