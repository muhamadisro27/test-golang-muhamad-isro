package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductCategory struct {
	ID   uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Name string    `json:"name"`
}

func (pc *ProductCategory) BeforeCreate(tx *gorm.DB) (err error) {
	pc.ID = uuid.New()
	return
}
