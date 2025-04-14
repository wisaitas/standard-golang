package models

import (
	"github.com/google/uuid"
	"github.com/wisaitas/standard-golang/pkg"
)

type District struct {
	pkg.BaseModel
	NameTH     string
	NameEN     string
	ProvinceID uuid.UUID
}
