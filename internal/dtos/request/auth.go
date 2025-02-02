package request

import "github.com/wisaitas/standard-golang/internal/models"

type LoginRequest struct {
	Username string `json:"username" validate:"required,min=3,max=255"`
	Password string `json:"password" validate:"required,min=8"`
}

type RegisterRequest struct {
	Username        string `json:"username" validate:"required,min=3,max=255"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=8"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
}

func (r *RegisterRequest) ToModel() models.User {
	return models.User{
		Username: r.Username,
		Email:    r.Email,
		Password: r.Password,
	}
}
