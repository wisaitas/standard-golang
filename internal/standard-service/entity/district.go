package entity

import (
	"github.com/google/uuid"
	entitySharePackage "github.com/wisaitas/share-pkg/db/entity"
)

type District struct {
	entitySharePackage.Entity

	NameTH string `gorm:"column:name_th;"`
	NameEN string `gorm:"column:name_en;"`

	ProvinceID uuid.UUID `gorm:"column:province_id;"`
}
