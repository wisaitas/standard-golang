package request

import (
	"time"

	"github.com/wisaitas/standard-golang/internal/standard-service/entity"
)

type CreateUserRequest struct {
	Username        string `json:"username" validate:"required,min=3,max=255"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=8"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
}

func (r *CreateUserRequest) RequestToEntity() entity.User {
	return entity.User{
		Username: r.Username,
		Email:    r.Email,
		Password: r.Password,
	}
}

type UpdateUserRequest struct {
	FirstName *string    `json:"first_name" validate:"omitempty,min=3,max=255"`
	LastName  *string    `json:"last_name" validate:"omitempty,min=3,max=255"`
	BirthDate *time.Time `json:"birth_date" validate:"omitempty"`
	Email     *string    `json:"email" validate:"omitempty,email"`
}
