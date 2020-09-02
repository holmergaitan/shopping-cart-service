package dao

import (
	"github.com/google/uuid"
	"shopping-cart-service/src/main/domain"
)

type GenericDao interface {
	Get(u uuid.UUID) *domain.Cart

	CreateOrUpdate(cart *domain.Cart) *domain.Cart

	RetrieveAll() []domain.Cart
}