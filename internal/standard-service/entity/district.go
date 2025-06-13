package entity

import (
	"github.com/google/uuid"
	"github.com/wisaitas/standard-golang/pkg"
)

type District struct {
	pkg.BaseEntity

	NameTH string `gorm:"column:name_th;"`
	NameEN string `gorm:"column:name_en;"`

	ProvinceID uuid.UUID `gorm:"column:province_id;"`

	Province *Province `gorm:"foreignKey:ProvinceID;references:ID"`
}
