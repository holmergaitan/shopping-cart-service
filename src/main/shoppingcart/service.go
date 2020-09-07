package shoppingcart

type Service struct{
	CartDao    GenericDao
	OrderDao   OrderDao
	ItemsCache CacheInterface
}

func (s *Service) CreateCart(c *Cart) *Cart {
	return s.CartDao.Create(c)
}

func (s *Service) GetCart(id string) *Cart {
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

func (s *Service) GetOrderByCartId (cartId string) *[]Order {
	return s.OrderDao.GetByCartId(cartId)
}