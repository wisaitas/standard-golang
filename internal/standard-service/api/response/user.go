package response

import (
	"time"

	"github.com/wisaitas/share-pkg/response"
	"github.com/wisaitas/standard-golang/internal/standard-service/entity"
)

type CreateUserResponse struct {
	response.EntityResponse
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (r *CreateUserResponse) EntityToResponse(entity entity.User) CreateUserResponse {
	r.EntityResponse = entity.EntityToResponse()
	r.Username = entity.Username
	r.Email = entity.Email

	return *r
}

type GetUsersResponse struct {
	response.EntityResponse
	Username  string            `json:"username"`
	Email     string            `json:"email"`
	FirstName string            `json:"first_name"`
	LastName  string            `json:"last_name"`
	BirthDate time.Time         `json:"birth_date"`
	Addresses []AddressResponse `json:"addresses"`
}

func (r *GetUsersResponse) EntityToResponse(entity entity.User) GetUsersResponse {
	r.EntityResponse = entity.EntityToResponse()
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
	response.EntityResponse
	Username  string            `json:"username"`
	Email     string            `json:"email"`
	FirstName string            `json:"first_name"`
	LastName  string            `json:"last_name"`
	BirthDate time.Time         `json:"birth_date"`
	Addresses []AddressResponse `json:"addresses"`
}

func (r *UpdateUserResponse) EntityToResponse(entity entity.User) UpdateUserResponse {
	r.EntityResponse = entity.EntityToResponse()
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
