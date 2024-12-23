package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"size:100;not null"`
	Phone       string `gorm:"uniqueIndex;size:20;not null"`
	CountryCode string `gorm:"size:5;not null"`
}
