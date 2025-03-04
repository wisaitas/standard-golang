package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uuid.UUID       `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	CreatedAt time.Time       `gorm:"type:timestamp;not null;default:now()"`
	UpdatedAt time.Time       `gorm:"type:timestamp;not null;default:now()"`
	DeletedAt *gorm.DeletedAt `gorm:"type:timestamp;default:null"`
}
