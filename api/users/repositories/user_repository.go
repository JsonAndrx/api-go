package repositories

import (
	"api-rest/config/database"
	"api-rest/api/users/models"
)


func CreateUserRepository(user *models.UserModel) (bool, error) {
	db := database.ConectDb()
	user_result := db.Create(&user)

	if user_result.Error != nil {
		return false, user_result.Error
	} 

	return true, nil
}