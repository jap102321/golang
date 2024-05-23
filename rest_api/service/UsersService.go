package service

import (
	"net/http"

	"example.com/rest-api/models"
	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func Signup(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request sign up data.",
		})
		return
	}

	err = user.SaveUser()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not save user data.",
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "User created succesfully",
	})
}

func Login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request login data.",
		})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Credentials invalid",
		})
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not generate token",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "login succesfully",
		"token":   token,
	})
}
