package initial

import (
	"github.com/wisaitas/standard-golang/internal/validates"
	"github.com/wisaitas/standard-golang/pkg"
)

type Validate struct {
	UserValidate        validates.UserValidate
	AuthValidate        validates.AuthValidate
	ProvinceValidate    validates.ProvinceValidate
	DistrictValidate    validates.DistrictValidate
	SubDistrictValidate validates.SubDistrictValidate
}

func NewValidate(validator pkg.ValidatorUtil) *Validate {
	return &Validate{
		UserValidate:        *validates.NewUserValidate(validator),
		AuthValidate:        *validates.NewAuthValidate(validator),
		ProvinceValidate:    *validates.NewProvinceValidate(validator),
		DistrictValidate:    *validates.NewDistrictValidate(validator),
		SubDistrictValidate: *validates.NewSubDistrictValidate(validator),
	}
}
