package requests

import (
	"time"

	"github.com/wisaitas/standard-golang/internal/standard-service/models"
)

type LoginRequest struct {
	Username string `json:"username" validate:"required,min=3,max=255" example:"john_doe"`
	Password string `json:"password" validate:"required,min=8" example:"Str0ngP@ssw0rd"`
}

type RegisterRequest struct {
	Username        string           `json:"username" validate:"required,min=3,max=255" example:"john_doe"`
	Email           string           `json:"email" validate:"required,email" example:"john.doe@example.com"`
	FirstName       string           `json:"first_name" validate:"required,min=3,max=255" example:"John"`
	LastName        string           `json:"last_name" validate:"required,min=3,max=255" example:"Doe"`
	BirthDate       time.Time        `json:"birth_date" validate:"required" example:"1990-01-01T00:00:00Z"`
	Password        string           `json:"password" validate:"required,min=8" example:"Str0ngP@ssw0rd"`
	ConfirmPassword string           `json:"confirm_password" validate:"required,eqfield=Password" example:"Str0ngP@ssw0rd"`
	Addresses       []AddressRequest `json:"addresses" validate:"dive"`
}

func (r *RegisterRequest) ReqToModel() models.User {
	addresses := []models.Address{}
	for _, address := range r.Addresses {
		addresses = append(addresses, address.ReqToModel())
	}

	return models.User{
		Username:  r.Username,
		Email:     r.Email,
		Password:  r.Password,
		FirstName: r.FirstName,
		LastName:  r.LastName,
		BirthDate: r.BirthDate,
		Addresses: addresses,
	}
}
