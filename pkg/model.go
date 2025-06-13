package pkg

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseEntity struct {
	ID        uuid.UUID       `gorm:"column:id;type:uuid;primaryKey;default:uuid_generate_v4()"`
	Version   int             `gorm:"column:version;type:integer;not null;default:0"`
	CreatedAt time.Time       `gorm:"column:created_at;type:timestamp;not null;default:now()"`
	CreatedBy *uuid.UUID      `gorm:"column:created_by;type:uuid"`
	UpdatedAt time.Time       `gorm:"column:updated_at;type:timestamp;not null;default:now()"`
	UpdatedBy *uuid.UUID      `gorm:"column:updated_by;type:uuid"`
	DeletedAt *gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp"`
	DeletedBy *uuid.UUID      `gorm:"column:deleted_by;type:uuid"`
}

func (r *BaseEntity) BeforeUpdate(tx *gorm.DB) (err error) {
	r.Version++
	return
}
