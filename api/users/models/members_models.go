package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Members struct {
	gorm.Model
	UserID            uint
	User              User `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
	TypeMembershipID  uint
	TypeMembership    TypeMembership `gorm:"foreignkey:TypeMembershipID;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
	ExpiredMembership time.Time      `gorm:"type:datetime;not null"`
}

type TypeMembership struct {
	gorm.Model
	TypeMembership     string `gorm:"type:varchar(12);not null"`
	DayMembership      int    `gorm:"type:int;not null"`
	AmountNotification int    `gorm:"type:int;not null"`
}
