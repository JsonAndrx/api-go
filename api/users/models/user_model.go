package models

import (
	"github.com/jinzhu/gorm"
)

type UserModel struct {
    gorm.Model
    FirstName      string `gorm:"type:varchar(12);not null" json:"first_name"`
    LastName       string `gorm:"type:varchar(12);not null" json:"last_name"`
    Email          string `gorm:"type:varchar(50);not null" json:"email"`
    Password       string `gorm:"type:varchar(100);not null" json:"password"`
    Country        string `gorm:"type:varchar(50);not null" json:"country"`
    TypeMembership int    `gorm:"type:int;not null;default:0" json:"type_membership"`
    CodeMembership string `gorm:"type:varchar(12)" json:"code_membership"`
    Role           string `gorm:"type:varchar(12)" json:"role"`
    IsActive       bool   `gorm:"default:true" json:"is_active"`
}
