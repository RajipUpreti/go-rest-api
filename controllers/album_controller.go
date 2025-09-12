package controllers

import (
	"albums-rest-api/models"
	"net/http"
	"strconv"

	"log/slog"

	"github.com/gin-gonic/gin"
)

func GetAlbums(c *gin.Context) {
	slog.Info("GetAlbums route hit")
	artist := c.Query("artist")
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, err1 := strconv.Atoi(pageStr)
	limit, err2 := strconv.Atoi(limitStr)

	if err1 != nil || err2 != nil || page <= 0 || limit <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid pagination parameters"})
		return
	}

	albums := models.GetAlbumsPaginatedFiltered(artist, page, limit)
	c.JSON(http.StatusOK, albums)
}

func GetAlbumByID(c *gin.Context) {
	slog.Info("GetAlbumByID route hit")
	id := c.Param("id")
	album, found := models.GetAlbumByID(id)

	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "album not found"})
		return
	}

	c.JSON(http.StatusOK, album)
}

func DeleteAlbumByID(c *gin.Context) {
	slog.Info("DeleteAlbumByID route hit")
	id := c.Param("id")

	deleted := models.DeleteAlbumByID(id)
	if !deleted {
		c.JSON(http.StatusNotFound, gin.H{"error": "album not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "album deleted"})
}

func CreateAlbum(c *gin.Context) {
	slog.Info("CreateAlbum route hit")
	var newAlbum models.Album

	// Bind JSON body into newAlbum struct
	if err := c.ShouldBindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	addedAlbum := models.AddAlbum(newAlbum)
	c.JSON(http.StatusCreated, addedAlbum)
}
