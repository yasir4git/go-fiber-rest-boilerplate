package models

import (
	"time"

	"gorm.io/gorm"
)

type CoreModel struct {
	ID        string    `json:"id" gorm:"type:uuid;primaryKey"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

type BaseModel struct {
	CoreModel
	CreatedBy *string        `gorm:"null;type:uuid"`
	UpdatedBy *string        `gorm:"null;type:uuid"`
	Updater   *User          `gorm:"foreignKey:UpdatedBy"`
	Creater   *User          `gorm:"foreignKey:CreatedBy"`
	DeletedAt gorm.DeletedAt `gorm:"uniqueIndex"`
}
