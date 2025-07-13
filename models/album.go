package models

import "github.com/google/uuid"

type Album struct {
	ID     string  `json:"id"`                            // must be present
	Title  string  `json:"title" binding:"required"`      // must be present
	Artist string  `json:"artist" binding:"required"`     // must be present
	Price  float64 `json:"price" binding:"required,gt=0"` // must be > 0
}

// Simulated in-memory data
var albums = []Album{
	{ID: uuid.New().String(), Title: "The Go Gospels", Artist: "Go Lang", Price: 9.99},
	{ID: uuid.New().String(), Title: "Gin and JSON", Artist: "Gin Gonic", Price: 12.50},
	{ID: uuid.New().String(), Title: "REST Rhapsody", Artist: "APIson", Price: 8.75},
}

func GetAlbums() []Album {
	return albums
}

func GetAlbumByID(id string) (*Album, bool) {
	for _, a := range albums {
		if a.ID == id {
			return &a, true
		}
	}
	return nil, false
}

func DeleteAlbumByID(id string) bool {
	for i, a := range albums {
		if a.ID == id {
			albums = append(albums[:i], albums[i+1:]...)
			return true
		}
	}
	return false
}

func AddAlbum(newAlbum Album) Album {
	newAlbum.ID = uuid.New().String() // assign UUID here
	albums = append(albums, newAlbum)
	return newAlbum
}
