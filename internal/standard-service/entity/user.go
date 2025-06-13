package entity

import (
	"time"

	"github.com/wisaitas/standard-golang/pkg"
)

type User struct {
	pkg.BaseEntity

	Username  string    `gorm:"column:username;"`
	FirstName string    `gorm:"column:first_name;"`
	LastName  string    `gorm:"column:last_name;"`
	BirthDate time.Time `gorm:"column:birth_date;"`
	Email     string    `gorm:"column:email;"`
	Password  string    `gorm:"column:password;"`

	Addresses []Address `gorm:"foreignKey:UserID"`
}
