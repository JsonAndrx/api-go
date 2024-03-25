package database

import (

	// "time"

	"api-rest/api/users/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func ConectDb() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@(db_maria:3306)/notihub")
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.UserModel{})

	return db
}
