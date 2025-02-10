package response

import "github.com/wisaitas/standard-golang/internal/models"

type CreateUserResponse struct {
	BaseResponse
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (r *CreateUserResponse) ToResponse(user models.User) CreateUserResponse {
	r.ID = user.ID
	r.CreatedAt = user.CreatedAt
	r.UpdatedAt = user.UpdatedAt
	r.Username = user.Username
	r.Email = user.Email

	return *r
}

type GetUsersResponse struct {
	BaseResponse
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (r *GetUsersResponse) ToResponse(users models.User) GetUsersResponse {
	r.ID = users.ID
	r.CreatedAt = users.CreatedAt
	r.UpdatedAt = users.UpdatedAt
	r.Username = users.Username
	r.Email = users.Email

	return *r
}
