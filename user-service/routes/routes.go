package routes

import (
	"user-service/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	r.POST("/register", controllers.CreateNewUser)
	r.GET("/", controllers.GetUsers)
	r.GET("/:id", controllers.GetUsersByID)
	r.POST("/login", controllers.LoginUser)

}
func HealthCheckRoutes(r *gin.Engine) {
	r.GET("/health", controllers.HealthCheck)
}
