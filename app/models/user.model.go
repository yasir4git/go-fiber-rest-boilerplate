package models

type User struct {
	BaseModel
	Name        string  `gorm:"size:100;not null"`
	Password    string  `gorm:"size:100;not null"`
	Phone       string  `gorm:"uniqueIndex;size:20;not null"`
	Email       *string `gorm:"uniqueIndex;size:100;null"`
	CountryCode string  `gorm:"size:5;not null"`
	IsActive    bool    `gorm:"default:false"`
}
