// Code generated by SQL-to-Model generator.
package models

import (
	"time"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SubDistrict struct {
	ID        uuid.UUID       `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Version   int             `gorm:"type:integer;not null;default:0"`
	CreatedAt time.Time       `gorm:"type:timestamp;not null;default:now()"`
	CreatedBy *uuid.UUID      `gorm:"type:uuid"`
	UpdatedAt time.Time       `gorm:"type:timestamp;not null;default:now()"`
	UpdatedBy *uuid.UUID      `gorm:"type:uuid"`
	DeletedAt *gorm.DeletedAt `gorm:"type:timestamp"`

	NameTh string `gorm:"type:varchar(100);not null"`
	NameEn string `gorm:"type:varchar(100);not null"`
	PostalCode string `gorm:"type:varchar(10);not null"`

	DistrictID uuid.UUID `gorm:"type:uuid;column:district_id"`

	District *District `gorm:"foreignKey:DistrictID;references:ID"`
	Addresses []Address `gorm:"foreignKey:SubDistrictID"`
}

func (r *SubDistrict) BeforeUpdate(tx *gorm.DB) (err error) {
	r.Version++
	return
}
