package repositories

import (
	"api-rest/api/users/models"
	"api-rest/config/database"
	"github.com/jinzhu/gorm"
)


func CreateMemberRepository(member *models.Members) (uint, error) {
	db := database.ConectDb()
	memberResult := db.Create(&member)

	if memberResult.Error != nil {
		return 0, memberResult.Error
	}

	return member.ID, nil
}

func GetTypeMembershipsRepository() ([]models.TypeMembership, error) {
	db := database.ConectDb()
	var typeMemberships []models.TypeMembership
	typeMembershipsResult := db.Find(&typeMemberships)

	if typeMembershipsResult.Error != nil {
		return nil, typeMembershipsResult.Error
	}

	return typeMemberships, nil
}

func GetTypeMembershipsByIdRepository(id uint) (int, error) {
    db := database.ConectDb()
    var typeMembership models.TypeMembership
    result := db.Where("id = ?", id).First(&typeMembership)
    if result.Error != nil {
        if result.Error == gorm.ErrRecordNotFound {
            return 0, nil
        }
        return 0, result.Error
    }
    return typeMembership.DayMembership, nil
}