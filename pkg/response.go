package pkg

import (
	"time"

	"github.com/google/uuid"
)

type SuccessResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type BaseResponse struct {
	ID        uuid.UUID  `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	CreatedBy *uuid.UUID `json:"created_by"`
	UpdatedAt time.Time  `json:"updated_at"`
	UpdatedBy *uuid.UUID `json:"updated_by"`
}

func (r *BaseResponse) ModelToResponse(model BaseModel) BaseResponse {
	r.ID = model.ID
	r.CreatedAt = model.CreatedAt
	r.UpdatedAt = model.UpdatedAt
	r.CreatedBy = model.CreatedBy
	r.UpdatedBy = model.UpdatedBy

	return *r
}
