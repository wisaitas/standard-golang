package models

import (
	"github.com/google/uuid"
	"github.com/wisaitas/standard-golang/pkg"
)

type SubDistrict struct {
	pkg.BaseModel
	NameTH     string
	NameEN     string
	DistrictID uuid.UUID
	PostalCode string
}
