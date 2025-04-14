package models

import (
	"github.com/google/uuid"
	"github.com/wisaitas/standard-golang/pkg"
)

type Address struct {
	pkg.BaseModel
	ProvinceID    uuid.UUID
	DistrictID    uuid.UUID
	SubDistrictID uuid.UUID
	Address       *string

	UserID uuid.UUID

	Province    *Province
	District    *District
	SubDistrict *SubDistrict
}
