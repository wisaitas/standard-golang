package entity

import (
	"github.com/wisaitas/standard-golang/pkg"
)

type Province struct {
	pkg.BaseEntity

	NameTH string `gorm:"column:name_th;"`
	NameEN string `gorm:"column:name_en;"`
}
