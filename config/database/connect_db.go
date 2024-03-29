package database

import (

	// "time"

	UserModel "api-rest/api/users/models"
	ClientModel "api-rest/api/clients/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func ConectDb() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@(db_maria:3306)/notihub?parseTime=true")
	if err != nil {
		panic(err)
	}

	db.Model(&UserModel.Members{}).AddForeignKey("type_membership_id", "type_memberships(id)", "CASCADE", "CASCADE")
	db.Model(&UserModel.Members{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(&UserModel.User{}).AddForeignKey("membership_id", "members(id)", "CASCADE", "CASCADE")
	db.Model(&ClientModel.Client{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.AutoMigrate(&UserModel.User{}, &UserModel.TypeMembership{}, &UserModel.Members{}, &ClientModel.Client{})

	seedTypeMemberships(db)

	return db
}

func seedTypeMemberships(db *gorm.DB) {
	var count int
    db.Model(&UserModel.TypeMembership{}).Count(&count)
    if count > 0 {
        return
    }

	typeMemberships := []UserModel.TypeMembership{
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
