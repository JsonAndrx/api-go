package database

import (

	// "time"

	"api-rest/api/users/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func ConectDb() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@(db_maria:3306)/notihub?parseTime=true")
	if err != nil {
		panic(err)
	}

	db.Model(&models.Members{}).AddForeignKey("type_membership_id", "type_memberships(id)", "CASCADE", "CASCADE")
	db.Model(&models.Members{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(&models.User{}).AddForeignKey("membership_id", "members(id)", "CASCADE", "CASCADE")
	db.AutoMigrate(&models.User{}, &models.TypeMembership{}, &models.Members{})

	seedTypeMemberships(db)

	return db
}

func seedTypeMemberships(db *gorm.DB) {
	var count int
    db.Model(&models.TypeMembership{}).Count(&count)
    if count > 0 {
        return
    }

	typeMemberships := []models.TypeMembership{
		{TypeMembership: "Free", DayMembership: 15, AmountNotification: 25},
		{TypeMembership: "Basic", DayMembership: 30, AmountNotification: 150},
		{TypeMembership: "Pro", DayMembership: 30, AmountNotification: 300},
	}

	for _, tm := range typeMemberships {
		if db.NewRecord(tm) {
			db.Create(&tm)
		}
	}
}
