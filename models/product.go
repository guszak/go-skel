package models

import "github.com/jinzhu/gorm"

// Product Model
type Product struct {
	gorm.Model
	CompanyID   uint
	Description string `gorm:"not null" form:"description" json:"description"`
	Unity       string `gorm:"not null" form:"unity" json:"unity"`
}
