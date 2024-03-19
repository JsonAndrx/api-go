package repository

import (
	"api-rest/helper"
	"api-rest/users/models"
	"api-rest/config/database"
)

func GetAllUsers() ([]models.User) {
	db := database.ConectDb()
	var users []models.User
	err := db.Select(&users, "SELECT * FROM accounts")
	if err != nil {
		helper.ErrorPanic(err)
	}  
	

	return users
}
