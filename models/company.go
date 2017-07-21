package models

import "github.com/jinzhu/gorm"

// Company Model
type Company struct {
	gorm.Model
	Name  string `gorm:"not null" form:"name" json:"name"`
	Token string `gorm:"not null" form:"token" json:"token"`
}
