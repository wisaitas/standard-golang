package responses

import (
	"time"

	"github.com/wisaitas/standard-golang/internal/standard-service/models"
	"github.com/wisaitas/standard-golang/pkg"
)

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (r *LoginResponse) ToResponse(accessToken, refreshToken string) LoginResponse {
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

func (r *RegisterResponse) ToResponse(user models.User) RegisterResponse {
	r.ID = user.ID
	r.CreatedAt = user.CreatedAt
	r.UpdatedAt = user.UpdatedAt
	r.Username = user.Username
	r.FirstName = user.FirstName
	r.LastName = user.LastName
	r.BirthDate = user.BirthDate
	r.Email = user.Email

	for _, address := range user.Addresses {
		addressResponse := AddressResponse{}
		r.Addresses = append(r.Addresses, addressResponse.ModelToResponse(address))
	}

	if len(r.Addresses) == 0 {
		r.Addresses = []AddressResponse{}
	}

	return *r
}
