package entity

import "github.com/wisaitas/share-pkg/db/entity"

type Province struct {
	entity.Entity

	NameTH string `gorm:"column:name_th;"`
	NameEN string `gorm:"column:name_en;"`
}
