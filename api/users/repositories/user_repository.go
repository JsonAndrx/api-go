package repositories

import (
	"api-rest/api/users/models"
	"api-rest/config/database"
)

func CreateUserRepository(user *models.UserModel) (bool, error) {
	db := database.ConectDb()
	userResult := db.Create(&user)

	if userResult.Error != nil {
		return false, userResult.Error
	}

	return true, nil
}
