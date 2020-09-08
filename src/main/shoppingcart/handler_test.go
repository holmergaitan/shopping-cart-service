package shoppingcart

import (
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
)

var s Config

func TestMain(m *testing.M) {
	dsn := "root:test@123@tcp(127.0.0.1:3306)/shopping_cart_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	_ = db.AutoMigrate(&Item{})
	_ = db.AutoMigrate(&Order{})
	_ = db.AutoMigrate(&Cart{})
	var cartDao = &CartDbRepository{Database: db}
	var itemDetailDao = &OrderDbRepository{Database: db}
	var cache = &ArticleDb{Database: db}
	cache.Load()

	var service = Service{CartDao: cartDao, OrderDao:itemDetailDao, ItemsCache: cache}
	var mapper = Mapper{Service: service}
	s = New(service, mapper)
	code := m.Run()
	os.Exit(code)
}

func TestApi_AddItem(t *testing.T) {
	req, err := http.NewRequest("GET", "/carts", nil)
	if err != nil {
		t.Fatal(err)
	}

	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func TestApi_CreateCart(t *testing.T) {
	body := []byte(`{"description":"Test description"}`)
	request, _ := http.NewRequest("POST", "/carts", bytes.NewBuffer(body))
	response := executeRequest(request)
	checkResponseCode(t, http.StatusCreated, response.Code)

	var cart Cart
	json.Unmarshal(response.Body.Bytes(), &cart)

	if cart.Description != "Test description" {
		t.Errorf("Expected user name to be 'test user'. Got '%v'", cart.Description)
	}

	if cart.Id == uuid.Nil {
		t.Errorf("Expected product ID to be '1'. Got '%v'", cart.Id)
	}
}

func TestApi_DeleteItem(t *testing.T) {
	var cart = getCartContent()
	CreateOrUpdate(&cart)
	req, err := http.NewRequest("DELETE", "/carts/"+ cart.Id.String() + "/items/2", nil)

	if err != nil {
		t.Fatal(err)
	}

	response := executeRequest(req)
	checkResponseCode(t, http.StatusCreated, response.Code)
	var cartFromApi Cart
	json.Unmarshal(response.Body.Bytes(), &cartFromApi)
	if cart.Id != cartFromApi.Id {
		t.Errorf("Expected a cart with id %s. Got %s", cart.Id, cartFromApi.Id)
	}
	if len(cartFromApi.Orders) != 1 {
		t.Errorf("Expected an item list with size %s. Got %s", "1", strconv.Itoa(len(cartFromApi.Orders)))
	}
}

func TestApi_DeleteItems(t *testing.T) {
	var cart = getCartContent()
	CreateOrUpdate(&cart)
	//req, err := http.NewRequest("DELETE", "/carts/"+ cart.Id.String() + "/items/2", nil)

}

func TestApi_GetArticles(t *testing.T) {

}

func TestApi_GetCarts(t *testing.T) {
	var cart = Cart{Description: "Test description"}
	CreateOrUpdate(&cart)
	req, err := http.NewRequest("GET", "/carts", nil)

	if err != nil {
		t.Fatal(err)
	}

	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
	var cartList[]Cart
	json.Unmarshal(response.Body.Bytes(), &cartList)
	if body := response.Body.String(); body == "[]" {
		t.Errorf("Expected an array. Got %s", body)
	}
	if len(cartList) != 1 {
		t.Errorf("Expected an array with %s elements. Got %s", "1", strconv.Itoa(len(cartList)))
	}
}

func TestApi_GetItems(t *testing.T) {

}

func Test_buildErrorResponse(t *testing.T) {

}

func Test_buildResponse(t *testing.T) {

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

//func getCartContent() Cart {
//	var items = map[string]Item{
//		"1":{
//			Id:"1",
//			Title:"Banana",
//			Price:"2.50",
//			Quantity: 1,
//		},
//		"2":{
//			Id:"2",
//			Title:"Apple",
//			Price:"3.20",
//			Quantity: 1,
//		},
//	}
//	var cart = Cart{
//		Description: "Test cart",
//		Orders:      items,
//	}
//
//	return cart
//}


