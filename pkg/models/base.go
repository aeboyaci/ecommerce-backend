package models

import (
	"github.com/google/uuid"
	"time"
)

type Base struct {
	ID        uuid.UUID `json:"-" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
