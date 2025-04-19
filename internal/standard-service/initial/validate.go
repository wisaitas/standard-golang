package initial

import (
	"github.com/wisaitas/standard-golang/internal/standard-service/validates"
)

type validate struct {
	userValidate        validates.UserValidate
	authValidate        validates.AuthValidate
	provinceValidate    validates.ProvinceValidate
	districtValidate    validates.DistrictValidate
	subDistrictValidate validates.SubDistrictValidate
}

func newValidate(util *util) *validate {
	return &validate{
		userValidate:        validates.NewUserValidate(util.validatorUtil),
		authValidate:        validates.NewAuthValidate(util.validatorUtil),
		provinceValidate:    validates.NewProvinceValidate(util.validatorUtil),
		districtValidate:    validates.NewDistrictValidate(util.validatorUtil),
		subDistrictValidate: validates.NewSubDistrictValidate(util.validatorUtil),
	}
}
