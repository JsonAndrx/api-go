package services

import (
	"fmt"
	"net/http"

	"api-rest/api/users/repositories"
	"api-rest/api/users/models"
	"api-rest/api/users/types"
	"api-rest/config/database"
	"api-rest/api/utils/location"
	"api-rest/api/utils/response"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

func GetUsersService(c *gin.Context) {
	db := database.ConectDb()
	response, err := fmt.Printf("Get all users from database: %v", db)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": response,
	})
}

func CreateUserService(c *gin.Context) {
	var request types.UserCreateRequest

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid request", err.Error()))
		return
	}

	cotunry_response, err := location.GetLocation(c.ClientIP())
	if err != nil {
		panic(err)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(http.StatusInternalServerError, response.ErrorResponse("Error while hashing the password", err.Error()))
        return
    }

	user := models.UserModel{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Password:  string(hashedPassword),
		Country:   cotunry_response,
		TypeMembership: 0,
		CodeMembership: "test",
		Role: "user",
		IsActive: true,
	}

	create_user, err := repositories.CreateUserRepository(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse("Error while creating the user", err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse("User created successfully", create_user))
}
