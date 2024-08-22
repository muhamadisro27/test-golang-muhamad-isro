package web

import "github.com/google/uuid"

type ProductResponse struct {
	ID          uuid.UUID   `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Category    interface{} `json:"category"`
}
