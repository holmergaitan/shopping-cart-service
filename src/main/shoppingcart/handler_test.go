package shoppingcart

import (
	"bytes"
	"encoding/json"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var s Config
var id string
var service Service

func TestMain(m *testing.M) {
	dsn := "root:test@123@tcp(127.0.0.1:3306)/shopping_cart_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	_ = db.AutoMigrate(&Item{})
	_ = db.AutoMigrate(&Order{})
	_ = db.AutoMigrate(&Cart{})
	var cartRepository = &CartDbRepository{Database: db}
	var orderRepository = &OrderDbRepository{Database: db}
	var cache = &ArticleDb{Database: db}
	cache.LoadAllItems()

	service = Service{CartDao: cartRepository, OrderDao: orderRepository, ItemsCache: cache}
	var mapper = Mapper{}
	s = New(service, mapper)
	code := m.Run()
	os.Exit(code)
}

func TestApi_GetCarts(t *testing.T) {
	req, err := http.NewRequest("GET", "/carts", nil)
	if err != nil {
		t.Fatal(err)
	}

	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
	if body := response.Body.String(); body == "[]" {
		t.Errorf("Expected an array with content. Got %s", body)
	}
}

func TestApi_Articles(t *testing.T) {
	req, err := http.NewRequest("GET", "/articles", nil)
	if err != nil {
		t.Fatal(err)
	}

	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
	if body := response.Body.String(); body == "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}


func TestApi_SaveCart(t *testing.T) {
	body := []byte(`{"description":"Test description"}`)
	request, _ := http.NewRequest("POST", "/carts", bytes.NewBuffer(body))
	response := executeRequest(request)
	checkResponseCode(t, http.StatusCreated, response.Code)
	var cart Cart
	json.Unmarshal(response.Body.Bytes(), &cart)

	if cart.Description != "Test description" {
		t.Errorf("Expected user name to be 'test user'. Got '%v'", cart.Description)
	}

	if cart.ID ==  "" {
		t.Errorf("Expected product ID to be '%v'. Got '%v'", cart.ID, "")
	}
}

func TestApi_AddItem(t *testing.T) {
	cartId := LoadData()
	body := []byte(`{
				"id": "4",
				"title": "Noodles",
				"price": "23.50"
			}`)

	request, _ := http.NewRequest("PUT", "/carts/"+ cartId + "/items", bytes.NewBuffer(body))
	response := executeRequest(request)
	checkResponseCode(t, http.StatusCreated, response.Code)
	order, err := service.GetOrderByCartAndItem(cartId, "4")
	if err != nil{
		t.Errorf("Unexpected error: %v", err.Error())
	}

	if order == nil{
		t.Errorf("Order should not be empty")
	}

	item := order.Item
	if item.Price != "23.50"{
		t.Errorf("Expected %v price. Got %v", "2.50", item.Price)
	}

	if item.Title != "Noodles"{
		t.Errorf("Expected %v title. Got %v", "Noodles", item.Title)
	}
}

func TestApi_DeleteItems(t *testing.T) {
	cartId := LoadData()
	order := Order{
		Quantity: 1,
		ItemId: "5",
		CartId: cartId,
	}
	_, err := service.CreateOrder(&order)
	if err != nil{
		t.Errorf("Unexpected error: %v", err.Error())
	}
	request, err := http.NewRequest("DELETE", "/carts/" + cartId + "/items/5", nil)
	if err != nil {
		t.Fatal(err)
	}

	response := executeRequest(request)
	checkResponseCode(t, http.StatusOK, response.Code)

	cart, err := service.GetCart(cartId)
	if cart.Orders != nil {
		t.Errorf("The order %v was found for cart with id %v ", cart.Orders, cart.ID)
	}
}

func TestApi_DeleteItem(t *testing.T) {
	cartId := LoadData()
	order := Order{
		Quantity: 1,
		ItemId: "5",
		CartId: cartId,
	}
	_, err := service.CreateOrder(&order)
	if err != nil{
		t.Errorf("Unexpected error: %v", err.Error())
	}
	request, err := http.NewRequest("DELETE", "/carts/" + cartId + "/items", nil)
	if err != nil {
		t.Fatal(err)
	}

	response := executeRequest(request)
	checkResponseCode(t, http.StatusOK, response.Code)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	s.Router().ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func LoadData() string{
	cart := Cart{Description: "Test cart"}
	service.CreateCart(&cart)
	return cart.ID
}

