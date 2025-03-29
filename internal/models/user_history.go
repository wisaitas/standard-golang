package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/wisaitas/standard-golang/pkg"
)

type UserHistory struct {
	pkg.BaseModel
	Action       string    `gorm:"not null"`
	OldVersion   int       `gorm:"not null"`
	OldFirstName string    `gorm:"not null"`
	OldLastName  string    `gorm:"not null"`
	OldBirthDate time.Time `gorm:"not null"`
	OldPassword  string    `gorm:"not null"`
	OldEmail     string    `gorm:"not null"`

	UserID uuid.UUID `gorm:"not null;index"`
}
