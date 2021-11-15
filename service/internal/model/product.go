package model

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Name      string    `gorm:"type:character(256)" json:"name"`
	CreateAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt
}

// TableName Rename
func (Product) TableName() string {
	return "product"
}
