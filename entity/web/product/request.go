package web

import "github.com/google/uuid"

type ProductCreateRequest struct {
	Name        string    `validate:"required,max=100,min=1"`
	Description string    `validate:"required,max=100,min=1"`
	CategoryId  uuid.UUID `validate:"required"`
}
