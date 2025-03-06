package responses

import (
	"time"

	"github.com/wisaitas/standard-golang/internal/models"
)

type CreateUserResponse struct {
	BaseResponse
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (r *CreateUserResponse) ModelToResponse(user models.User) CreateUserResponse {
	r.ID = user.ID
	r.CreatedAt = user.CreatedAt
	r.UpdatedAt = user.UpdatedAt
	r.Username = user.Username
	r.Email = user.Email

	return *r
}

type GetUsersResponse struct {
	BaseResponse
	Username  string            `json:"username"`
	Email     string            `json:"email"`
	FirstName string            `json:"first_name"`
	LastName  string            `json:"last_name"`
	BirthDate time.Time         `json:"birth_date"`
	Addresses []AddressResponse `json:"addresses"`
}

func (r *GetUsersResponse) ModelToResponse(user models.User) GetUsersResponse {
	r.ID = user.ID
	r.CreatedAt = user.CreatedAt
	r.UpdatedAt = user.UpdatedAt
	r.Username = user.Username
	r.Email = user.Email
	r.FirstName = user.FirstName
	r.LastName = user.LastName
	r.BirthDate = user.BirthDate

	for _, address := range user.Addresses {
		addressResponse := AddressResponse{}
		r.Addresses = append(r.Addresses, addressResponse.ModelToResponse(address))
	}

	if len(r.Addresses) == 0 {
		r.Addresses = []AddressResponse{}
	}

	return *r
}

type UpdateUserResponse struct {
	BaseResponse
	Username  string            `json:"username"`
	Email     string            `json:"email"`
	FirstName string            `json:"first_name"`
	LastName  string            `json:"last_name"`
	BirthDate time.Time         `json:"birth_date"`
	Addresses []AddressResponse `json:"addresses"`
}

func (r *UpdateUserResponse) ModelToResponse(user models.User) UpdateUserResponse {
	r.ID = user.ID
	r.CreatedAt = user.CreatedAt
	r.UpdatedAt = user.UpdatedAt
	r.Username = user.Username
	r.Email = user.Email
	r.FirstName = user.FirstName
	r.LastName = user.LastName
	r.BirthDate = user.BirthDate

	for _, address := range user.Addresses {
		addressResponse := AddressResponse{}
		r.Addresses = append(r.Addresses, addressResponse.ModelToResponse(address))
	}

	if len(r.Addresses) == 0 {
		r.Addresses = []AddressResponse{}
	}

	return *r
}
