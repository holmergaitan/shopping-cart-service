package domain

import "github.com/google/uuid"

type Cart struct {
	Id          uuid.UUID 		`json:"id"`
	Description string    		`json:"description,omitempty" validate:"nil=false > empty=false" `
	Items       map[string]Item `json:"item"`
}
