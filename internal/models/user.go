package models

import (
	"time"

	"github.com/wisaitas/standard-golang/pkg"
)

type User struct {
	pkg.BaseModel
	Username  string    `gorm:"not null;unique"`
	FirstName string    `gorm:"not null"`
	LastName  string    `gorm:"not null"`
	BirthDate time.Time `gorm:"not null"`
	Email     string    `gorm:"not null;unique"`
	Password  string    `gorm:"not null"`

	Addresses []Address `gorm:"foreignKey:UserID"`
}
