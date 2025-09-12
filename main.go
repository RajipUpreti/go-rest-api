package main

import (
	"albums-rest-api/config"
	"albums-rest-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.ConnectDatabase()
	// config.DB.AutoMigrate(&models.Album{})

	// models.SeedAlbumsIfEmpty() // ✅ run once if DB is empty

	routes.AlbumRoutes(r)
	routes.HealthCheckRoutes(r)
	r.Run()
}
