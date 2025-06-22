package entity

import (
	"github.com/google/uuid"
	"github.com/wisaitas/share-pkg/db/entity"
)

type Address struct {
	entity.Entity

	Address *string `gorm:"column:address;"`

	ProvinceID    uuid.UUID `gorm:"column:province_id;"`
	DistrictID    uuid.UUID `gorm:"column:district_id;"`
	SubDistrictID uuid.UUID `gorm:"column:sub_district_id;"`
	UserID        uuid.UUID `gorm:"column:user_id;"`

	Province    *Province    `gorm:"foreignKey:ProvinceID;references:ID"`
	District    *District    `gorm:"foreignKey:DistrictID;references:ID"`
	SubDistrict *SubDistrict `gorm:"foreignKey:SubDistrictID;references:ID"`
	User        *User        `gorm:"foreignKey:UserID;references:ID"`
}
