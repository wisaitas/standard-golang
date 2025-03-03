package responses

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
	Username  string            `json:"username"`
	Email     string            `json:"email"`
	Addresses []AddressResponse `json:"addresses"`
}

func (r *GetUsersResponse) ModelToResponse(users models.User) GetUsersResponse {
	r.ID = users.ID
	r.CreatedAt = users.CreatedAt
	r.UpdatedAt = users.UpdatedAt
	r.Username = users.Username
	r.Email = users.Email

	for _, address := range users.Addresses {
		addressResponse := AddressResponse{}
		r.Addresses = append(r.Addresses, addressResponse.ModelToResponse(address))
	}

	if len(r.Addresses) == 0 {
		r.Addresses = []AddressResponse{}
	}

	return *r
}
