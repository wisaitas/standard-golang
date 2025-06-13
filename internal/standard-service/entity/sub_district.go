package entity

import (
	"github.com/google/uuid"
	"github.com/wisaitas/standard-golang/pkg"
)

type SubDistrict struct {
	pkg.BaseEntity

	NameTH     string `gorm:"column:name_th;"`
	NameEN     string `gorm:"column:name_en;"`
	PostalCode string `gorm:"column:postal_code;"`

	DistrictID uuid.UUID `gorm:"column:district_id;"`

	District *District `gorm:"foreignKey:DistrictID;references:ID"`
}
