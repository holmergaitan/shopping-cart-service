package shoppingcart

type Mapper struct {
	Service Service
}

type ItemDto struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Price  string  `json:"price"`
}

type OrderDto struct{
	ID 			string		`json:"id"`
	Item        ItemDto	   `json:"item"`
	Quantity 	int			`json:"quantity"`
}

type CartDto struct {
	ID          string  	`json:"id"`
	Description string  	`json:"description"`
	Orders      []OrderDto 	`json:"orders"`
}

func (m *Mapper) ToCartDtoList(carts []Cart) *[]CartDto{
	cartDtoList := make([]CartDto, len(carts))
	for i, cart := range carts{
		cartDtoList[i] = *m.ToCartDto(cart)
	}
	return &cartDtoList
}

func (m *Mapper) ToCartDto(cart Cart) *CartDto{
	var orders = cart.Orders
	ordersDao := make([]OrderDto, len(orders))
	for i, order := range orders{
		ordersDao[i] = *m.ToOrderDto(order)
	}

	var cartDto = &CartDto{
		ID: cart.ID,
		Description: cart.Description,
		Orders: ordersDao,
	}

	return cartDto
}

func (m *Mapper) ToCarDomain(cartDto CartDto) *Cart{
	orders := make([]Order, 0)
	for _, orderDto := range cartDto.Orders{
		od := *m.ToOrderDomain(orderDto, cartDto.ID, orderDto.Item.ID)
		orders = append(orders, od)
	}
	cart := &Cart{
		ID:          cartDto.ID,
		Description: cartDto.Description,
		Orders:      orders,
	}

	return cart
}


func (m *Mapper) ToOrderDtoList(orders []Order) *[]OrderDto {
	orderDtoList := make([]OrderDto, len(orders))
	for i, cart := range orders{
		orderDtoList[i] = *m.ToOrderDto(cart)
	}
	return &orderDtoList
}

func (m *Mapper) ToOrderDto(order Order) *OrderDto {
	item := order.Item
	itemDto := m.ToItemDto(item)
	orderDto := &OrderDto{
		ID:       	order.ID,
		Item:    	*itemDto,
		Quantity: 	order.Quantity,
	}
	return orderDto
}

func (m *Mapper) ToOrderDomain(orderDto OrderDto, cartId string, itemId string) *Order {
	order := &Order{
		ID: orderDto.ID,
		CartId: cartId,
		ItemId: itemId,
		Quantity: orderDto.Quantity,
	}
	return order
}

func (m *Mapper) ToItemDto(item Item) *ItemDto {
	itemDto := &ItemDto{
		ID: item.ID,
		Price: item.Price,
		Title: item.Title,
	}

	return itemDto
}

func (m *Mapper) ToItemDtoList(items []Item) *[]ItemDto{
	itemDtoList := make([]ItemDto, len(items))
	for i, item := range items{
		itemDtoList[i] = *m.ToItemDto(item)
	}
	return &itemDtoList
}

func (m *Mapper) ToItemDomain(itemDto ItemDto) *Item {
	item := &Item{
		ID: itemDto.ID,
		Title: itemDto.Title,
		Price: itemDto.Price,
	}
	return item
}
