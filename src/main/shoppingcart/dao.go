package shoppingcart

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"sync"
)

type GenericDao interface {
	Get(string) *Cart

	Create(cart *Cart) *Cart

	Update(cart *Cart) *Cart

	GetAll() *[]Cart
}

type OrderDao interface {
	Create(order *Order) *Order

	Update(order *Order) *Order

	GetByCartId(orderId string) *[]Order

	GetByCartAndItem(cartId string, itemId string) *Order
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

func (d *CartDbDao) Get(id string) *Cart {
	var cart Cart = Cart{}
	d.Database.Where("id = ?", id).First(&cart)
	return &cart
}

func (d *CartDbDao) Create(cart *Cart) *Cart {
	var id, _ = uuid.NewRandom()
	cart.ID = id.String()
	d.Database.Create(cart)
	return cart
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

//func (d *CartMapDao) GetByCartAndItem(cartId string, itemId string) *Order {
//	return nil
//}
