package models

import (
	"time"
)

type User struct {
	BaseModel
	Username  string    `gorm:"not null;unique"`
	FirstName string    `gorm:"not null"`
	LastName  string    `gorm:"not null"`
	BirthDate time.Time `gorm:"not null"`
	Email     string    `gorm:"not null;unique"`
	Password  string    `gorm:"not null"`

	Addresses []Address `gorm:"foreignKey:UserID"`
}
