package controllers

import (
	"go-rest-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAlbums(c *gin.Context) {
	albums := models.GetAlbums()
	c.JSON(http.StatusOK, albums)
}

func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")
	album, found := models.GetAlbumByID(id)

	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "album not found"})
		return
	}

	c.JSON(http.StatusOK, album)
}

func DeleteAlbumByID(c *gin.Context) {
	id := c.Param("id")

	deleted := models.DeleteAlbumByID(id)
	if !deleted {
		c.JSON(http.StatusNotFound, gin.H{"error": "album not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "album deleted"})
}

func CreateAlbum(c *gin.Context) {
	var newAlbum models.Album

	// Bind JSON body into newAlbum struct
	if err := c.ShouldBindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	addedAlbum := models.AddAlbum(newAlbum)
	c.JSON(http.StatusCreated, addedAlbum)
}
