package repositories

import (
	"api-rest/api/users/models"
	"api-rest/config/database"
	"github.com/jinzhu/gorm"

)

func CreateUserRepository(user *models.User) (uint, error) {
    db := database.ConectDb()
    userResult := db.Create(&user)

    if userResult.Error != nil {
        return 0, userResult.Error
    }

    return user.ID, nil
}

func GetUsersRepository() ([]models.User, error) {
	db := database.ConectDb()
	var users []models.User
	userResult := db.Find(&users)

	if userResult.Error != nil {
		return nil, userResult.Error
	}

	return users, nil
}

func GetUserByEmailRepository(email string) (bool, error) {
    db := database.ConectDb()
    var user models.User
    result := db.Select("email").Where("email = ?", email).First(&user)
    if result.Error != nil {
        if result.Error == gorm.ErrRecordNotFound {
            return false, nil
        }
        return false, result.Error
    }
    return true, nil
}


func GetUserByIdRepository(id uint) (bool, error) {
	db := database.ConectDb()
	var user models.User
	result := db.Where("id = ?", id).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, result.Error
	}
	return true, nil
}

func UpdateUserMembershipRepository(id uint, idMembership uint) (bool, error) {
	db := database.ConectDb()
	result := db.Model(&models.User{}).Where("id = ?", id).Update("membership_id", idMembership)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}