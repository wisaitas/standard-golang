package response

import "github.com/wisaitas/standard-golang/internal/models"

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
	BaseResponse
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (r *RegisterResponse) ToResponse(user models.User) RegisterResponse {
	r.ID = user.ID
	r.CreatedAt = user.CreatedAt
	r.UpdatedAt = user.UpdatedAt
	r.Username = user.Username
	r.Email = user.Email

	return *r
}
