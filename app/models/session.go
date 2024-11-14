package models

import "gorm.io/gorm"

type Session struct {
	gorm.Model
	UserID       uint
	RefreshToken string `gorm:"unique;not null"`
}
