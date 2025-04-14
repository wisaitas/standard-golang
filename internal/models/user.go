package models

import (
	"time"

	"github.com/wisaitas/standard-golang/pkg"
)

type User struct {
	pkg.BaseModel
	Username  string
	FirstName string
	LastName  string
	BirthDate time.Time
	Email     string
	Password  string

	Addresses []Address
}
