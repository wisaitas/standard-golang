package initial

import (
	"github.com/wisaitas/standard-golang/internal/validates"
)

type Validates struct {
	UserValidate        validates.UserValidate
	AuthValidate        validates.AuthValidate
	ProvinceValidate    validates.ProvinceValidate
	DistrictValidate    validates.DistrictValidate
	SubDistrictValidate validates.SubDistrictValidate
}

func initializeValidates() *Validates {
	return &Validates{
		UserValidate:        *validates.NewUserValidate(),
		AuthValidate:        *validates.NewAuthValidate(),
		ProvinceValidate:    *validates.NewProvinceValidate(),
		DistrictValidate:    *validates.NewDistrictValidate(),
		SubDistrictValidate: *validates.NewSubDistrictValidate(),
	}
}
