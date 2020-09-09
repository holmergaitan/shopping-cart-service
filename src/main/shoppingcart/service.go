package shoppingcart

type Service struct{
	CartDao    	CartRepository
	OrderDao   	OrderRepository
	ItemsCache 	CacheInterface
}

func (s *Service) CreateCart(c *Cart) (*Cart, error) {
	return s.CartDao.SaveCart(c)
}

func (s *Service) GetCart(id string) (*Cart, error) {
	return s.CartDao.GetCart(id)
}

func (s *Service) GetCarts() (*[]Cart, error) {
	carts, err := s.CartDao.GetAllCarts()
	var cartList = make([]Cart, len(*carts))
	for i, cart := range *carts{
		orders, _ := s.GetOrderByCartId(cart.ID)
		cartList[i] = Cart{
			ID: cart.ID,
			Description: cart.Description,
			Orders: orders,
		}
	}
	return &cartList, err
}

func (s *Service) CreateOrder(detail *Order) (*Order, error) {
	return s.OrderDao.CreateOrder(detail)
}

func (s *Service) GetOrderByCartAndItem(cartId string, itemId string) (*Order, error) {
	return s.OrderDao.GetOrderByCartAndItemId(cartId, itemId)
}

func (s *Service) UpdateOrder(i *Order) (*Order, error) {
	return s.OrderDao.UpdateOrder(i)
}

func (s *Service) DeleteOrder(i *Order) error {
	return s.OrderDao.DeleteOrder(i)
}

func (s *Service) GetOrderByCartId(cartId string) (*[]Order, error) {
	return s.OrderDao.GetOrdersByCartId(cartId)
}

func (s *Service) DeleteOrdersByCart(cartId string) error {
	return s.OrderDao.DeleteOrderByCartId(cartId)
}

func (s *Service) GetItem(id string) (*Item, error) {
	return s.ItemsCache.GetItemById(id)
}