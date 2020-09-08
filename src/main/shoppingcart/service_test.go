package shoppingcart

import (
	"reflect"
	"testing"
)

func TestService_CreateCart(t *testing.T) {
	type fields struct {
		CartDao    CartRepository
		OrderDao   OrderRepository
		ItemsCache CacheInterface
	}
	type args struct {
		c *Cart
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Cart
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				CartDao:    tt.fields.CartDao,
				OrderDao:   tt.fields.OrderDao,
				ItemsCache: tt.fields.ItemsCache,
			}
			got, err := s.CreateCart(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("SaveCart() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SaveCart() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_CreateOrder(t *testing.T) {
	type fields struct {
		CartDao    CartRepository
		OrderDao   OrderRepository
		ItemsCache CacheInterface
	}
	type args struct {
		detail *Order
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Order
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				CartDao:    tt.fields.CartDao,
				OrderDao:   tt.fields.OrderDao,
				ItemsCache: tt.fields.ItemsCache,
			}
			if got, _ := s.CreateOrder(tt.args.detail); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_DeleteOrder(t *testing.T) {
	type fields struct {
		CartDao    CartRepository
		OrderDao   OrderRepository
		ItemsCache CacheInterface
	}
	type args struct {
		i *Order
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				CartDao:    tt.fields.CartDao,
				OrderDao:   tt.fields.OrderDao,
				ItemsCache: tt.fields.ItemsCache,
			}
			if got := s.DeleteOrder(tt.args.i); got != tt.want {
				t.Errorf("DeleteOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_DeleteOrdersByCart(t *testing.T) {
	type fields struct {
		CartDao    CartRepository
		OrderDao   OrderRepository
		ItemsCache CacheInterface
	}
	type args struct {
		cartId string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				CartDao:    tt.fields.CartDao,
				OrderDao:   tt.fields.OrderDao,
				ItemsCache: tt.fields.ItemsCache,
			}
			if got := s.DeleteOrdersByCart(tt.args.cartId); got != tt.want {
				t.Errorf("DeleteOrdersByCart() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetCart(t *testing.T) {
	type fields struct {
		CartDao    CartRepository
		OrderDao   OrderRepository
		ItemsCache CacheInterface
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Cart
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				CartDao:    tt.fields.CartDao,
				OrderDao:   tt.fields.OrderDao,
				ItemsCache: tt.fields.ItemsCache,
			}
			got, err := s.GetCart(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCart() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCart() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetCarts(t *testing.T) {
	type fields struct {
		CartDao    CartRepository
		OrderDao   OrderRepository
		ItemsCache CacheInterface
	}
	tests := []struct {
		name   string
		fields fields
		want   *[]Cart
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				CartDao:    tt.fields.CartDao,
				OrderDao:   tt.fields.OrderDao,
				ItemsCache: tt.fields.ItemsCache,
			}
			if got := s.GetCarts(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCarts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetOrderByCartAndItem(t *testing.T) {
	type fields struct {
		CartDao    CartRepository
		OrderDao   OrderRepository
		ItemsCache CacheInterface
	}
	type args struct {
		cartId string
		itemId string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Order
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				CartDao:    tt.fields.CartDao,
				OrderDao:   tt.fields.OrderDao,
				ItemsCache: tt.fields.ItemsCache,
			}
			if got, _ := s.GetOrderByCartAndItem(tt.args.cartId, tt.args.itemId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOrderByCartAndItemId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetOrderByCartId(t *testing.T) {
	type fields struct {
		CartDao    CartRepository
		OrderDao   OrderRepository
		ItemsCache CacheInterface
	}
	type args struct {
		cartId string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *[]Order
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				CartDao:    tt.fields.CartDao,
				OrderDao:   tt.fields.OrderDao,
				ItemsCache: tt.fields.ItemsCache,
			}
			if got, _ := s.GetOrderByCartId(tt.args.cartId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOrderByCartId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_UpdateOrder(t *testing.T) {
	type fields struct {
		CartDao    CartRepository
		OrderDao   OrderRepository
		ItemsCache CacheInterface
	}
	type args struct {
		i *Order
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Order
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				CartDao:    tt.fields.CartDao,
				OrderDao:   tt.fields.OrderDao,
				ItemsCache: tt.fields.ItemsCache,
			}
			if got := s.UpdateOrder(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
