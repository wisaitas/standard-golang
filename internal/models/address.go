package models

import (
	"github.com/google/uuid"
)

type Address struct {
	BaseModel
	ProvinceID    int
	DistrictID    int
	SubDistrictID int
	Address       *string

	UserID uuid.UUID

	Province    *Province
	District    *District
	SubDistrict *SubDistrict
}
