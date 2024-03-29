package models

import (
	UserModel "api-rest/api/users/models"
	"time"

	"github.com/jinzhu/gorm"
)

type Client struct {
	gorm.Model
	UserID      uint
	User        UserModel.User `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
	FirtsName   string         `gorm:"type:varchar(12);not null" json:"firts_name"`
	LastName    string         `gorm:"type:varchar(12);not null" json:"last_name"`
	Phone       string         `gorm:"type:varchar(12);not null" json:"phone"`
	Date        time.Time      `gorm:"type:datetime;not null" json:"date"`
	Description string         `gorm:"type:varchar(500);not null" json:"description"`
}
