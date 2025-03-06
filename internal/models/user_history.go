package models

import (
	"time"

	"github.com/google/uuid"
)

type UserHistory struct {
	BaseModel
	Action    string    `gorm:"not null"`
	UserID    uuid.UUID `gorm:"not null"`
	Version   int       `gorm:"not null"`
	FirstName string    `gorm:"not null"`
	LastName  string    `gorm:"not null"`
	BirthDate time.Time `gorm:"not null"`
}
