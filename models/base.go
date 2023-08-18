package models

import (
	"github.com/google/uuid"
	"time"
)

type Model struct {
	ID			uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	CreatedAt	time.Time
	UpdatedAt	time.Time
}