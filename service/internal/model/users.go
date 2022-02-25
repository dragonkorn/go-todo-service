package model

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	MobileNo string `gorm:"type:character(10)" json:"mobileNo"`
	Username string `gorm:"type:character(256)" json:"username"`
}

// TableName Rename
func (Users) TableName() string {
	return "users"
}
