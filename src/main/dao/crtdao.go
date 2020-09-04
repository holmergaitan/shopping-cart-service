package dao

import (
	"github.com/google/uuid"
	"shopping-cart-service/src/main/domain"
	"sync"
)

type GenericDao interface {
	Get(u uuid.UUID) *domain.Cart

	CreateOrUpdate(cart *domain.Cart) *domain.Cart

	RetrieveAll() []domain.Cart
}

type CartDao struct {
	Carts map[uuid.UUID]domain.Cart
	mux sync.Mutex
}

func (d *CartDao) Get(id uuid.UUID) *domain.Cart {
	elem, _ := d.Carts[id]
	return &elem
}

func (d *CartDao) CreateOrUpdate(cart *domain.Cart) *domain.Cart {
	if cart.Id == uuid.Nil {
		cart.Id, _ = uuid.NewRandom()
	}
	d.Carts[cart.Id] = *cart
	return cart
}

func (d *CartDao) RetrieveAll() []domain.Cart {
	var carts []domain.Cart
	for _, value := range d.Carts {
		carts = append(carts, value)
	}
	return carts
}
