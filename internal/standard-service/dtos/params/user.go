package params

import "github.com/google/uuid"

type UserParams struct {
	ID uuid.UUID `param:"id" validate:"required,uuid"`
}
