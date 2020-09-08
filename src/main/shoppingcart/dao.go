package shoppingcart

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"sync"
)

type CartDao interface {
	Get(string) (*Cart, error)

	Create(cart *Cart) (*Cart, error)

	Update(cart *Cart) *Cart

	GetAll() *[]Cart
}

type OrderDao interface {
	Create(order *Order) *Order

	Update(order *Order) *Order

	GetByCartId(orderId string) *[]Order

	GetByCartAndItem(cartId string, itemId string) *Order

	DeleteByCartId(id string) bool

	Delete(order *Order) bool
}

type CartDbDao struct{
	Database *gorm.DB
}

type OrderDbDao struct{
	Database *gorm.DB
}

type CartMapDao struct {
	Carts map[string]Cart
	mux sync.Mutex
}

func (d *CartDbDao) Get(id string) (*Cart, error) {
	var cart Cart = Cart{}
	result := d.Database.Where("id = ?", id).First(&cart)
	if result.Error != nil {
		return nil, result.Error
	}

	return &cart, nil
}

func (d *CartDbDao) Create(cart *Cart) (*Cart, error) {
	var id, _ = uuid.NewRandom()
	cart.ID = id.String()
	result := d.Database.Create(cart)
	if result.Error != nil {
		return nil, result.Error
	}

	return cart, nil
}

func (d *CartDbDao) Update(cart *Cart) *Cart {
	var cartDb Cart
	d.Database.Model(&cartDb).Updates(cart)
	return cart
}

func (d *CartDbDao) GetAll() *[]Cart {
	var carts = make([]Cart, 0)
	d.Database.Preload("Orders").Find(&carts)
	return &carts
}

func (o *OrderDbDao) Create(itemDetail *Order) *Order {
	var id, _ = uuid.NewRandom()
	itemDetail.ID = id.String()
	o.Database.Create(itemDetail)
	return itemDetail
}

func (o *OrderDbDao) Update(order *Order) *Order {
	var orderDb Order
	o.Database.Model(&orderDb).Where("id = ?", order.ID).Updates(order)
	return order
}

func (o *OrderDbDao) GetByCartId(cartId string) *[]Order {
	var orders = make([]Order, 0)
	o.Database.Where("cart_id = ?", cartId).Find(&orders)
	return &orders
}

func (o *OrderDbDao) GetByCartAndItem(cartId string, itemId string) *Order {
	var orderDb Order
	o.Database.Where("cart_id = ? and item_id = ? ", cartId, itemId).First(&orderDb)
	return &orderDb
}

func (o *OrderDbDao) DeleteByCartId(cartId string) bool {
	var count int64
	o.Database.Delete(Order{}, "cart_id = ?", cartId)
	o.Database.Table("orders").Where("cart_id = ?", cartId).Count(&count)

	if count != 0 {
		return false
	}

	return true
}

func (o *OrderDbDao) Delete(order *Order) bool{
	var count int64
	o.Database.Delete(&order)
	o.Database.Table("orders").Where("id = ?", order.ID).Count(&count)

	if count != 0 {
		return false
	}

	return true
}

func (d *CartMapDao) Create(cart *Cart) *Cart {
	return &Cart{}
}

func (d *CartMapDao) Update(cart *Cart) *Cart {
	if cart.ID == "" {
		id, _ := uuid.NewRandom()
		cart.ID = id.String()
	}
	d.Carts[cart.ID] = *cart
	return cart
}

func (d *CartMapDao) GetAll() *[]Cart {
	var carts = make([]Cart, 0)
	for _, value := range d.Carts {
		carts = append(carts, value)
	}
	return &carts
}

func (d *CartMapDao) Get(id string) *Cart {
	elem, _ := d.Carts[id]
	return &elem
}
