package entity

import (
	"github.com/google/uuid"
	"github.com/wisaitas/share-pkg/db/entity"
)

type SubDistrict struct {
	entity.Entity

	NameTH     string `gorm:"column:name_th;"`
	NameEN     string `gorm:"column:name_en;"`
	PostalCode string `gorm:"column:postal_code;"`

	DistrictID uuid.UUID `gorm:"column:district_id;"`

	District *District `gorm:"foreignKey:DistrictID;references:ID"`
}
