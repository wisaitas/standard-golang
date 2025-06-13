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

func (r *BaseResponse) EntityToResponse(entity BaseEntity) BaseResponse {
	r.ID = entity.ID
	r.CreatedAt = entity.CreatedAt
	r.UpdatedAt = entity.UpdatedAt
	r.CreatedBy = entity.CreatedBy
	r.UpdatedBy = entity.UpdatedBy

	return *r
}
