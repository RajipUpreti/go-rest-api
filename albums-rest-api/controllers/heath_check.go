package controllers

import (
	"net/http"
	"time"

	"log/slog"

	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	slog.Info("HealthCheck route hit")
	c.JSON(http.StatusOK, gin.H{
		"status":    "ok",
		"message":   "API is running",
		"version":   "1.0.0",
		"timestamp": time.Now().Format(time.RFC3339),
	})
}
