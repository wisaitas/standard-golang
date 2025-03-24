package pkg

import "github.com/go-playground/validator/v10"

type ValidatorUtil interface {
	Validate(v any) error
}

type validatorUtil struct {
	validate *validator.Validate
}

func NewValidatorUtil() ValidatorUtil {
	return &validatorUtil{validate: validator.New()}
}

func (r *validatorUtil) Validate(v any) error {
	return r.validate.Struct(v)
}
