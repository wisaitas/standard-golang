package request

import (
	"time"

	"github.com/wisaitas/standard-golang/internal/standard-service/entity"
)

type LoginRequest struct {
	Username string `json:"username" validate:"required,min=3,max=255" example:"john_doe"`
	Password string `json:"password" validate:"required,min=8" example:"Str0ngP@ssw0rd"`
}

type RegisterRequest struct {
	Username        string           `json:"username" validate:"required,min=3,max=255"`
	Email           string           `json:"email" validate:"required,email"`
	FirstName       string           `json:"first_name" validate:"required,min=3,max=255"`
	LastName        string           `json:"last_name" validate:"required,min=3,max=255"`
	BirthDate       time.Time        `json:"birth_date" validate:"required"`
	Password        string           `json:"password" validate:"required,min=8"`
	ConfirmPassword string           `json:"confirm_password" validate:"required,eqfield=Password"`
	Addresses       []AddressRequest `json:"addresses" validate:"dive"`
}

func (r *RegisterRequest) RequestToEntity() entity.User {
	addresses := []entity.Address{}
	for _, address := range r.Addresses {
		addresses = append(addresses, address.RequestToEntity())
	}

	return entity.User{
		Username:  r.Username,
		Email:     r.Email,
		Password:  r.Password,
		FirstName: r.FirstName,
		LastName:  r.LastName,
		BirthDate: r.BirthDate,
		Addresses: addresses,
	}
}
