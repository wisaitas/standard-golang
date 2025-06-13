package initial

import (
	validateInternal "github.com/wisaitas/standard-golang/internal/standard-service/validate"
)

type validate struct {
	userValidate        validateInternal.UserValidate
	authValidate        validateInternal.AuthValidate
	provinceValidate    validateInternal.ProvinceValidate
	districtValidate    validateInternal.DistrictValidate
	subDistrictValidate validateInternal.SubDistrictValidate
}

func newValidate(lib *lib) *validate {
	return &validate{
		userValidate:        validateInternal.NewUserValidate(lib.validator),
		authValidate:        validateInternal.NewAuthValidate(lib.validator),
		provinceValidate:    validateInternal.NewProvinceValidate(lib.validator),
		districtValidate:    validateInternal.NewDistrictValidate(lib.validator),
		subDistrictValidate: validateInternal.NewSubDistrictValidate(lib.validator),
	}
}
