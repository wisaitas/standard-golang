package request

import "github.com/wisaitas/standard-golang/internal/models"

type LoginRequest struct {
	Username string `json:"username" validate:"required,min=3,max=255"`
	Password string `json:"password" validate:"required,min=8"`
}

type RegisterRequest struct {
	Username        string                   `json:"username" validate:"required,min=3,max=255"`
	Email           string                   `json:"email" validate:"required,email"`
	Password        string                   `json:"password" validate:"required,min=8"`
	ConfirmPassword string                   `json:"confirm_password" validate:"required,eqfield=Password"`
	Addresses       []RegisterAddressRequest `json:"addresses" validate:"dive"`
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
		Addresses: addresses,
	}
}

type RegisterAddressRequest struct {
	ProvinceID    int     `json:"province_id" validate:"required"`
	DistrictID    int     `json:"district_id" validate:"required"`
	SubDistrictID int     `json:"sub_district_id" validate:"required"`
	Address       *string `json:"address"`
}

func (r *RegisterAddressRequest) ReqToModel() models.Address {
	return models.Address{
		ProvinceID:    r.ProvinceID,
		DistrictID:    r.DistrictID,
		SubDistrictID: r.SubDistrictID,
		Address:       r.Address,
	}
}
