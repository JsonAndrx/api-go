package services

import (
	"fmt"
	"net/http"

	"api-rest/api/users/models"
	"api-rest/api/users/repositories"
	"api-rest/api/users/types"
	"api-rest/api/utils/location"
	"api-rest/api/utils/response"
	"os"

	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

func GetUsersService(c *gin.Context) {
	// configuracion pa mi envio de mensajes
	os.Setenv("TWILIO_ACCOUNT_SID", "AC2917c7ec5db82e70ba3bfc0de6df4513")
	os.Setenv("TWILIO_AUTH_TOKEN", "442e6e282c5f555a0c1e91ea0f209e86")

	fmt.Println(os.Environ())

	client := twilio.NewRestClient()

	params := &api.CreateMessageParams{}
	params.SetFrom("whatsapp:+14155238886")
	params.SetBody("men ahora estoy cambiando el mensaje a ver si se actualiza sin tener que construir otra vez el contenedor :)")
	params.SetTo("whatsapp:+573008362662")

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		if resp.Sid != nil {
			fmt.Println(*resp.Sid)
		} else {
			fmt.Println(resp.Sid)
		}
	}
}

func CreateUserService(c *gin.Context) {
	var request types.UserCreateRequest

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid request", err.Error()))
		return
	}

	exists, err := repositories.GetUserByEmailRepository(request.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse("Error while checking the email", err.Error()))
		return
	}
	if exists {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Email already exists", nil))
		return
	}

	countryResponse, err := location.GetLocation(c.ClientIP())
	if err != nil {
		panic(err)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse("Error while hashing the password", err.Error()))
		return
	}

	user := models.User{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Password:  string(hashedPassword),
		Country:   countryResponse,
		Role:      "user",
		IsActive:  true,
	}

	createdUser, err := repositories.CreateUserRepository(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse("Error while creating the user", err.Error()))
		return
	}

	createMembership, err := CreateMembersService(createdUser, 1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse("Error while creating the membership", err.Error()))
		return
	}

	_, err = repositories.UpdateUserMembershipRepository(createdUser, createMembership)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse("Error while updating the user", err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse("User created successfully", createMembership))
}
