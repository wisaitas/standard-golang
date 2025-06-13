package param

import "github.com/google/uuid"

type UserParam struct {
	ID uuid.UUID `param:"id" validate:"required,uuid"`
}
