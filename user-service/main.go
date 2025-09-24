package main

import (
	"user-service/config"
	"user-service/models"
	"user-service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.ConnectDatabase()
	if err := models.Migrate(config.DB); err != nil {
		panic(err)
	}

	routes.UserRoutes(r)
	routes.HealthCheckRoutes(r)
	r.Run()
	config.CloseDatabase()
}
