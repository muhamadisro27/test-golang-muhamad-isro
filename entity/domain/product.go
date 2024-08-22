package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID          uuid.UUID       `gorm:"type:uuid;primary_key" json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	CategoryID  uuid.UUID       `gorm:"type:uuid" json:"category_id"`
	Category    ProductCategory `gorm:"foreignKey:category_id" json:"category"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
    p.ID = uuid.New()
    return
}
