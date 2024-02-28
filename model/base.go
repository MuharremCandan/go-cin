package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID        uuid.UUID      `gorm:"type:uuid;primary_key" json:"id" example:"f0a0a1e9-69bd-4bef-b8c6-4e8c0d3a1212"`
	CreatedAt time.Time      `json:"created_at" example:"2021-01-01T00:00:00Z"`
	UpdatedAt time.Time      `json:"updated_at" example:"2021-01-01T00:00:00Z"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index" swaggertype:"string" format:"date-time"`
}

func (g *Base) BeforeCreate(tx *gorm.DB) (err error) {
	g.ID = uuid.New()
	return nil
}
