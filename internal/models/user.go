package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	BaseModel
	Username  string    `gorm:"not null;unique"`
	FirstName string    `gorm:"not null"`
	LastName  string    `gorm:"not null"`
	BirthDate time.Time `gorm:"not null"`
	Email     string    `gorm:"not null;unique"`
	Password  string    `gorm:"not null"`
	Version   int       `gorm:"not null;default:0"`

	Addresses []Address `gorm:"foreignKey:UserID"`
}

func (r *User) BeforeUpdate(tx *gorm.DB) (err error) {
	r.Version++
	return
}
