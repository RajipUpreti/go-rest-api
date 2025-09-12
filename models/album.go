package models

import (
	"go-rest-api/config"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Album struct {
	ID     uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Title  string    `json:"title" binding:"required"`
	Artist string    `json:"artist" binding:"required"`
	Price  float64   `json:"price" binding:"required,gt=0"`
}

func SeedAlbumsIfEmpty() {
	var count int64
	config.DB.Model(&Album{}).Count(&count)
	if count == 0 {
		albums := []Album{
			{Title: "Initial Track", Artist: "System", Price: 10.0},
			{Title: "Reload Safe", Artist: "GORM", Price: 20.0},
		}
		config.DB.Create(&albums)
	}
}

func GetAlbums() []Album {
	var albums []Album
	config.DB.Find(&albums)
	return albums
}

func GetAlbumByID(id string) (*Album, bool) {
	var album Album

	if err := config.DB.Take(&album, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, false
		}
		panic(err)
	}
	return &album, true
}

func AddAlbum(newAlbum Album) Album {
	config.DB.Create(&newAlbum)
	return newAlbum
}

func GetAlbumsPaginatedFiltered(artist string, page, limit int) []Album {
	var albums []Album
	db := config.DB

	if artist != "" {
		db = db.Where("artist LIKE ?", "%"+artist+"%")
	}

	offset := (page - 1) * limit
	db.Limit(limit).Offset(offset).Find(&albums)

	return albums
}

func DeleteAlbumByID(id string) bool {
	var album Album
	result := config.DB.Delete(&album, id)
	return result.RowsAffected > 0
}
