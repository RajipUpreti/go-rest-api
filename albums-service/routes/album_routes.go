package routes

import (
	"albums-service/controllers"

	"github.com/gin-gonic/gin"
)

func AlbumRoutes(r *gin.Engine) {
	r.GET("/", controllers.GetAlbums)
	r.GET("/:id", controllers.GetAlbumByID)
	r.DELETE("/:id", controllers.DeleteAlbumByID)
	r.POST("/", controllers.CreateAlbum)
	// r.POST("/dapr", controllers.SaveAlbumToDaprState)

}
func HealthCheckRoutes(r *gin.Engine) {
	r.GET("/health", controllers.HealthCheck)
}
