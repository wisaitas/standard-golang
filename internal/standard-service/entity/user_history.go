package entity

import (
	"time"

	"github.com/wisaitas/standard-golang/pkg"
)

type UserHistory struct {
	pkg.BaseEntity

	Action       string    `gorm:"column:action;"`
	OldVersion   int       `gorm:"column:old_version;"`
	OldFirstName string    `gorm:"column:old_first_name;"`
	OldLastName  string    `gorm:"column:old_last_name;"`
	OldBirthDate time.Time `gorm:"column:old_birth_date;"`
	OldPassword  string    `gorm:"column:old_password;"`
	OldEmail     string    `gorm:"column:old_email;"`
}
