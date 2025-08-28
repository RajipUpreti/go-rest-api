package main

import (
	"go-rest-api/config"
	"go-rest-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.ConnectDatabase()
	// config.DB.AutoMigrate(&models.Album{})

	// models.SeedAlbumsIfEmpty() // ✅ run once if DB is empty

	routes.RegisterAlbumRoutes(r)
	r.Run()
}
