package service

import (
	"github.com/google/uuid"
	"shopping-cart-service/src/main/dao"
	"shopping-cart-service/src/main/domain"
)

var cartDao = dao.CartDao{Carts: make(map[uuid.UUID]domain.Cart)}

func CreateOrUpdate(c *domain.Cart) *domain.Cart {
	return cartDao.CreateOrUpdate(c)
}

func GetCart(id uuid.UUID) *domain.Cart{
	return cartDao.Get(id)
}

func RetrieveCarts()[]domain.Cart{
	return cartDao.RetrieveAll()
}