package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
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

func (a *Address) BeforeCreate(tx *gorm.DB) error {
	a.ID = uuid.New()

	return nil
}
