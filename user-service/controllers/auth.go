package controllers

import (
	"net/http"
	"strings"
	"user-service/models"
	"user-service/utils"

	"log/slog"

	"github.com/gin-gonic/gin"
)

func CreateNewUser(c *gin.Context) {
	slog.Info("CreateNewUser route hit")
	var newUserDetails models.UserDetails

	if err := c.ShouldBindJSON(&newUserDetails); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid input: "+err.Error())
		return
	}

	addedUserDetails, err := models.AddUserWithRole(newUserDetails)
	if err != nil {
		// Detect duplicate email error
		if strings.Contains(err.Error(), "duplicate key") {
			utils.RespondWithError(c, http.StatusConflict, "Email already exists")
			return
		}

		utils.RespondWithError(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	// Sanitize response
	addedUserDetails.Hash = ""

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    addedUserDetails,
	})
}

func GetUsers(c *gin.Context) {
	slog.Info("GetUsers route hit")
	users := models.GetUsers()
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    users,
	})
}

func GetUsersByID(c *gin.Context) {
	slog.Info("GetUsersByID route hit")
	id := c.Param("id")
	user, found := models.GetUsersByID(id)
	if !found {
		utils.RespondWithError(c, http.StatusNotFound, "User not found")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    user,
	})
}

func LoginUser(c *gin.Context) {

	slog.Info("LoginUser route hit")
	var loginDetails struct {
		EmailID  string ` json:"emailID" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&loginDetails); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid input: "+err.Error())
		return
	}
	user, err := models.LoginUser(loginDetails.EmailID, loginDetails.Password)
	if err != nil {
		utils.RespondWithError(c, http.StatusUnauthorized, "Invalid credentials")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    user,
	})
}
