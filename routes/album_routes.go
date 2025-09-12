package routes

import (
	"go-rest-api/controllers"

	"github.com/gin-gonic/gin"
)

func AlbumRoutes(r *gin.Engine) { 
	r.GET("/albums", controllers.GetAlbums)
	r.GET("/albums/:id", controllers.GetAlbumByID)
	r.DELETE("/albums/:id", controllers.DeleteAlbumByID)
	r.POST("/albums", controllers.CreateAlbum)
	// r.POST("/dapr/albums", controllers.SaveAlbumToDaprState)

}
func HealthCheckRoutes(r *gin.Engine) {
	r.GET("/health", controllers.HealthCheck)
}
