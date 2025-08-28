package routes

import (
	"go-rest-api/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterAlbumRoutes(r *gin.Engine) {
	r.GET("/albums", controllers.GetAlbums)
	r.GET("/albums/:id", controllers.GetAlbumByID)
	r.DELETE("/albums/:id", controllers.DeleteAlbumByID)
	r.POST("/albums", controllers.CreateAlbum)
	r.POST("/dapr/albums", controllers.SaveAlbumToDaprState)

}
