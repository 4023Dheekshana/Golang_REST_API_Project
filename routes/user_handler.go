package routes

import (
	"log"
	"net/http"

	"dheek.com/restapi/models"
	"dheek.com/restapi/utils"
	"github.com/gin-gonic/gin"
)

func signUp(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		log.Fatalf("Error binding context %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}
	err = user.Save()
	if err != nil {
		log.Fatalf("Error saving the users in database %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the user data."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "user data saved successfully."})
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		log.Fatalf("Error binding context %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data"})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		log.Fatalf("Error validating credential %v", err)
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate password"})
		return
	}
	token, err := utils.GenerateToken(user.Email, user.Id)
	if err != nil {
		log.Fatalf("Error generating a token %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not generate a token."})
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login Successful.", "token": token})
}
