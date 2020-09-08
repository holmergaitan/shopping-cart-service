package shoppingcart

type Service struct{
	CartDao    CartDao
	OrderDao   OrderDao
	ItemsCache CacheInterface
}

func (s *Service) CreateCart(c *Cart) (*Cart, error) {
	return s.CartDao.Create(c)
}

func (s *Service) GetCart(id string) (*Cart, error) {
	return s.CartDao.Get(id)
}

func (s *Service) GetCarts() *[]Cart {
	return s.CartDao.GetAll()
}

func (s *Service) CreateOrder (detail *Order) *Order {
	return s.OrderDao.Create(detail)
}

func (s *Service) GetOrderByCartAndItem(cartId string, itemId string) *Order {
	return s.OrderDao.GetByCartAndItem(cartId, itemId)
}

func (s *Service) UpdateOrder (i *Order) *Order {
	return s.OrderDao.Update(i)
}

func (s *Service) DeleteOrder (i *Order) bool {
	return s.OrderDao.Delete(i)
}

func (s *Service) GetOrderByCartId (cartId string) *[]Order {
	return s.OrderDao.GetByCartId(cartId)
}

func (s *Service) DeleteOrdersByCart(cartId string) bool {
	return s.OrderDao.DeleteByCartId(cartId)
}