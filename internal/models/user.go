package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	BaseModel
	Username string `gorm:"not null;unique"`
	Email    string `gorm:"not null;unique"`
	Password string `gorm:"not null"`

	Addresses []Address `gorm:"foreignKey:UserID"`
}

func (r *User) BeforeCreate(tx *gorm.DB) error {
	r.ID = uuid.New()

	return nil
}
