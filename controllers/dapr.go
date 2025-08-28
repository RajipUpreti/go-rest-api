package controllers

import (
	"go-rest-api/models"
	"go-rest-api/utils"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func SaveAlbumToDaprState(c *gin.Context) {
	var album models.Album
	if err := c.ShouldBindJSON(&album); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ✅ Assign UUID if no ID is provided
	if album.ID == "" {
		album.ID = uuid.New().String()
	}

	state := []map[string]interface{}{
		{
			"key":   album.ID,
			"value": album,
		},
	}

	daprURL := "http://localhost:3501/v1.0/state/statestore"
	resp, err := http.Post(daprURL, "application/json", utils.ToJSONReader(state))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)

	c.JSON(http.StatusCreated, gin.H{
		"message":       "Album saved to Dapr state",
		"id":            album.ID, // ✅ return ID to client
		"status_code":   resp.StatusCode,
		"dapr_response": string(bodyBytes),
	})
}
