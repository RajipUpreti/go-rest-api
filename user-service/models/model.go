package models

import (
	"log/slog"
	"user-service/config"
	"user-service/utils"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserDetails struct {
	ID       uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserName string    `json:"userName" binding:"required"`
	EmailID  string    `gorm:"uniqueIndex:uidx_users_email" json:"emailID" binding:"required,email"`

	Role string `json:"role" binding:"required,oneof=admin user manager"`
	Hash string `gorm:"column:password_hash"`

	Password   string `json:"password" binding:"required"`
	ApiKEY     string `gorm:"uniqueIndex:uidx_users_api_key" json:"ApiKEY" `
	ConsumerID string `gorm:"uniqueIndex:uidx_users_consumer_id" json:"ConsumerID" `
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&UserDetails{})
}

func GetUsers() []UserDetails {
	var users []UserDetails
	config.DB.Find(&users)
	return users
}
func GetUsersByID(id string) (*UserDetails, bool) {
	var user UserDetails

	if err := config.DB.Take(&user, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, false
		}
		panic(err)
	}
	return &user, true
}

func AddUserWithRole(newUserDetails UserDetails) (UserDetails, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUserDetails.Password), bcrypt.DefaultCost)
	if err != nil {
		return UserDetails{}, err
	}
	consumer, err := utils.CreateKongConsumer(newUserDetails.EmailID)
	if err != nil {
		return UserDetails{}, err
	}

	// // randomised APIkey
	kongAPIKEY, err := CreateAPIKey(consumer.ID)
	if err != nil {
		slog.Info("Failed to create API key")
		return UserDetails{}, err
	}
	// newUserDetails.ApiKEY = ""
	newUserDetails.ConsumerID = consumer.ID
	newUserDetails.Password = ""
	newUserDetails.Hash = ""
	newUserDetails.ApiKEY = kongAPIKEY.Key
	// Assign hash to DB column
	newUserDetails.Hash = string(hashedPassword)

	// Create user (excluding plain password)
	result := config.DB.Create(&newUserDetails)
	if result.Error != nil {
		return UserDetails{}, result.Error
	}

	// Do not return hash to client
	newUserDetails.Hash = ""
	newUserDetails.Password = ""
	return newUserDetails, nil
}

func LoginUser(email string, password string) (*UserDetails, error) {
	var user UserDetails
	if err := config.DB.Where("email_id = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Hash), []byte(password)); err != nil {
		return nil, err
	}

	return &user, nil

}
