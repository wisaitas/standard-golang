package models

import (
	"time"

	"github.com/wisaitas/standard-golang/pkg"
)

type UserHistory struct {
	pkg.BaseModel
	Action       string
	OldVersion   int
	OldFirstName string
	OldLastName  string
	OldBirthDate time.Time
	OldPassword  string
	OldEmail     string
}
