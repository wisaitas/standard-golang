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

func newValidate(sharePkg *sharePkg) *validate {
	return &validate{
		userValidate:        validateInternal.NewUserValidate(sharePkg.validator),
		authValidate:        validateInternal.NewAuthValidate(sharePkg.validator),
		provinceValidate:    validateInternal.NewProvinceValidate(sharePkg.validator),
		districtValidate:    validateInternal.NewDistrictValidate(sharePkg.validator),
		subDistrictValidate: validateInternal.NewSubDistrictValidate(sharePkg.validator),
	}
}
