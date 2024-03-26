package services

import (
	"time"
	"api-rest/api/users/models"
	"api-rest/api/users/repositories"
	// "github.com/gin-gonic/gin"
	// "api-rest/api/utils/response"
	// "net/http"
)

func CreateMembersService(idUser uint, idTypeMembership uint) (uint, error) {

	userMember, err := repositories.GetUserByIdRepository(idUser)
	if err != nil {
		return 0, err
	}
	if !userMember {
		return 0, nil
	}

	typeMemberships, err := repositories.GetTypeMembershipsByIdRepository(idTypeMembership)
	if err != nil {
		return 0, err
	}
	if typeMemberships == 0 {
		return 0, nil
	}

	memberUser := models.Members{
		UserID:          idUser,
		TypeMembershipID: idTypeMembership,
		ExpiredMembership:  time.Now().AddDate(0, 0, int(typeMemberships)),
	}

	createdMember, err := repositories.CreateMemberRepository(&memberUser)
	if err != nil {
		return 0, err
	}

	return createdMember, nil
}
