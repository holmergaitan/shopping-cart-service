package shoppingcart

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"sync"
)

type CartRepository interface {
	GetCart(string) (*Cart, error)

	SaveCart(cart *Cart) (*Cart, error)

	UpdateCart(cart *Cart) *Cart

	GetAllCarts() *[]Cart
}

type OrderRepository interface {
	CreateOrder(order *Order) (*Order, error)

	UpdateOrder(order *Order) *Order

	GetOrderByCartId(orderId string) (*[]Order, error)

	GetOrderByCartAndItemId(cartId string, itemId string) (*Order, error)

	DeleteOrderByCartId(cartId string) error

	DeleteOrder(order *Order) error
}

type CartDbRepository struct{
	Database *gorm.DB
}

type OrderDbRepository struct{
	Database *gorm.DB
}

type CartMapRepository struct {
	Carts map[string]Cart
	mux sync.Mutex
}

func (d *CartDbRepository) GetCart(id string) (*Cart, error) {
	var cart Cart = Cart{}
	result := d.Database.Where("id = ?", id).First(&cart)
	if result.Error != nil {
		return nil, result.Error
	}

	return &cart, nil
}

func (d *CartDbRepository) SaveCart(cart *Cart) (*Cart, error) {
	var id, _ = uuid.NewRandom()
	cart.ID = id.String()
	result := d.Database.Create(cart)
	if result.Error != nil {
		return nil, result.Error
	}

	return cart, nil
}

func (d *CartDbRepository) UpdateCart(cart *Cart) *Cart {
	var cartDb Cart
	d.Database.Model(&cartDb).Updates(cart)
	return cart
}

func (d *CartDbRepository) GetAllCarts() *[]Cart {
	var carts = make([]Cart, 0)
	d.Database.Preload("Orders").Find(&carts)
	return &carts
}

func (o *OrderDbRepository) CreateOrder(order *Order) (*Order, error) {
	var id, _ = uuid.NewRandom()
	order.ID = id.String()
	result := o.Database.Create(order)
	if result.Error != nil {
		return nil, result.Error
	}

	return order, nil
}

func (o *OrderDbRepository) UpdateOrder(order *Order) *Order {
	var orderDb Order
	o.Database.Model(&orderDb).Where("id = ?", order.ID).Updates(order)
	return order
}

func (o *OrderDbRepository) GetOrderByCartId(orderId string) (*[]Order, error) {
	var orders = make([]Order, 0)
	result := o.Database.Where("cart_id = ?", orderId).Find(&orders)
	return &orders, result.Error
}

func (o *OrderDbRepository) GetOrderByCartAndItemId(cartId string, itemId string) (*Order, error) {
	var orderDb Order
	result := o.Database.Where("cart_id = ? and item_id = ? ", cartId, itemId).First(&orderDb)
	return &orderDb, result.Error
}

func (o *OrderDbRepository) DeleteOrderByCartId(cartId string) error {
	result := o.Database.Delete(Order{}, "cart_id = ?", cartId)
	return result.Error
}

func (o *OrderDbRepository) DeleteOrder(order *Order) error {
	result := o.Database.Delete(&order)
	return result.Error
}

func (d *CartMapRepository) Create(cart *Cart) *Cart {
	return &Cart{}
}

func (d *CartMapRepository) Update(cart *Cart) *Cart {
	if cart.ID == "" {
		id, _ := uuid.NewRandom()
		cart.ID = id.String()
	}
	d.Carts[cart.ID] = *cart
	return cart
}

func (d *CartMapRepository) GetAll() *[]Cart {
	var carts = make([]Cart, 0)
	for _, value := range d.Carts {
		carts = append(carts, value)
	}
	return &carts
}

func (d *CartMapRepository) Get(id string) *Cart {
	elem, _ := d.Carts[id]
	return &elem
}
