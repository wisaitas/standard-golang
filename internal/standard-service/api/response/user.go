package response

import (
	"time"

	"github.com/wisaitas/standard-golang/internal/standard-service/entity"
	"github.com/wisaitas/standard-golang/pkg"
)

type CreateUserResponse struct {
	pkg.BaseResponse
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (r *CreateUserResponse) EntityToResponse(entity entity.User) CreateUserResponse {
	r.BaseResponse.EntityToResponse(entity.BaseEntity)
	r.Username = entity.Username
	r.Email = entity.Email

	return *r
}

type GetUsersResponse struct {
	pkg.BaseResponse
	Username  string            `json:"username"`
	Email     string            `json:"email"`
	FirstName string            `json:"first_name"`
	LastName  string            `json:"last_name"`
	BirthDate time.Time         `json:"birth_date"`
	Addresses []AddressResponse `json:"addresses"`
}

func (r *GetUsersResponse) EntityToResponse(entity entity.User) GetUsersResponse {
	r.BaseResponse.EntityToResponse(entity.BaseEntity)
	r.Username = entity.Username
	r.Email = entity.Email
	r.FirstName = entity.FirstName
	r.LastName = entity.LastName
	r.BirthDate = entity.BirthDate

	for _, address := range entity.Addresses {
		addressResponse := AddressResponse{}
		r.Addresses = append(r.Addresses, addressResponse.EntityToResponse(address))
	}

	if len(r.Addresses) == 0 {
		r.Addresses = []AddressResponse{}
	}

	return *r
}

type UpdateUserResponse struct {
	pkg.BaseResponse
	Username  string            `json:"username"`
	Email     string            `json:"email"`
	FirstName string            `json:"first_name"`
	LastName  string            `json:"last_name"`
	BirthDate time.Time         `json:"birth_date"`
	Addresses []AddressResponse `json:"addresses"`
}

func (r *UpdateUserResponse) EntityToResponse(entity entity.User) UpdateUserResponse {
	r.BaseResponse.EntityToResponse(entity.BaseEntity)
	r.Username = entity.Username
	r.Email = entity.Email
	r.FirstName = entity.FirstName
	r.LastName = entity.LastName
	r.BirthDate = entity.BirthDate

	for _, address := range entity.Addresses {
		addressResponse := AddressResponse{}
		r.Addresses = append(r.Addresses, addressResponse.EntityToResponse(address))
	}

	if len(r.Addresses) == 0 {
		r.Addresses = []AddressResponse{}
	}

	return *r
}
