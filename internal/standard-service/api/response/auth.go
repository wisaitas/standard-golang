package response

import (
	"time"

	"github.com/wisaitas/standard-golang/internal/standard-service/entity"
	"github.com/wisaitas/standard-golang/pkg"
)

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (r *LoginResponse) EntityToResponse(accessToken, refreshToken string) LoginResponse {
	r.AccessToken = accessToken
	r.RefreshToken = refreshToken

	return *r
}

type RegisterResponse struct {
	pkg.BaseResponse
	Username  string            `json:"username"`
	FirstName string            `json:"first_name"`
	LastName  string            `json:"last_name"`
	BirthDate time.Time         `json:"birth_date"`
	Email     string            `json:"email"`
	Addresses []AddressResponse `json:"addresses"`
}

func (r *RegisterResponse) EntityToResponse(entity entity.User) RegisterResponse {
	r.BaseResponse.EntityToResponse(entity.BaseEntity)
	r.Username = entity.Username
	r.FirstName = entity.FirstName
	r.LastName = entity.LastName
	r.BirthDate = entity.BirthDate
	r.Email = entity.Email

	for _, address := range entity.Addresses {
		addressResponse := AddressResponse{}
		r.Addresses = append(r.Addresses, addressResponse.EntityToResponse(address))
	}

	if len(r.Addresses) == 0 {
		r.Addresses = []AddressResponse{}
	}

	return *r
}
