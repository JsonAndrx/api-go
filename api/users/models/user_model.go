package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	FirstName    string                `gorm:"type:varchar(12);not null" json:"first_name"`
	LastName     string                `gorm:"type:varchar(12);not null" json:"last_name"`
	Email        string                `gorm:"type:varchar(50);not null" json:"email"`
	Password     string                `gorm:"type:varchar(100);not null" json:"password"`
	Country      string                `gorm:"type:varchar(50);not null" json:"country"`
	MembershipID *uint                 `gorm:"null" json:"membership_id"`
	Membership   TypeMembership `gorm:"foreignkey:MembershipID;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
	Role         string                `gorm:"type:varchar(12)" json:"role"`
	IsActive     bool                  `gorm:"default:true" json:"is_active"`
}
