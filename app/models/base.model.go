package models

import (
	"time"

	"github.com/google/uuid"
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

// Hook BeforeCreate - set UUID and timestamps
func (m *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	// Set UUID if not already set
	if m.ID == "" {
		m.ID = uuid.New().String()
	}

	ctx := tx.Statement.Context
	if ctx != nil {
		if user, ok := ctx.Value("user").(*User); ok {
			m.CreatedBy = &user.ID
			m.UpdatedBy = &user.ID
		}
	}

	// Set CreatedAt and UpdatedAt
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

// Hook BeforeUpdate - update timestamps
func (m *BaseModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = time.Now()
	return nil
}
