package main

import (
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"net/http/httptest"
	"os"
	"shopping-cart-service/src/main/config"
	"shopping-cart-service/src/main/domain"
	"shopping-cart-service/src/main/service"
	"strconv"
	"testing"
)

var s config.Config

func TestMain(m *testing.M) {
	s = config.New()
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

	var cart domain.Cart
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
	service.CreateOrUpdate(&cart)
	req, err := http.NewRequest("DELETE", "/carts/"+ cart.Id.String() + "/items/2", nil)

	if err != nil {
		t.Fatal(err)
	}

	response := executeRequest(req)
	checkResponseCode(t, http.StatusCreated, response.Code)
	var cartFromApi domain.Cart
	json.Unmarshal(response.Body.Bytes(), &cartFromApi)
	if cart.Id != cartFromApi.Id {
		t.Errorf("Expected a cart with id %s. Got %s", cart.Id, cartFromApi.Id)
	}
	if len(cartFromApi.Items) != 1 {
		t.Errorf("Expected an item list with size %s. Got %s", "1", strconv.Itoa(len(cartFromApi.Items)))
	}
}

func TestApi_DeleteItems(t *testing.T) {
	var cart = getCartContent()
	service.CreateOrUpdate(&cart)
	//type fields struct {
	//	ApiRouter http.Handler
	//}
	//type args struct {
	//	w http.ResponseWriter
	//	r *http.Request
	//}
	//tests := []struct {
	//	name   string
	//	fields fields
	//	args   args
	//}{
	//	// TODO: Add test cases.
	//}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		a := &Api{
	//			ApiRouter: tt.fields.ApiRouter,
	//		}
	//		print(a)
	//	})
	//}
}

func TestApi_GetArticles(t *testing.T) {

}

func TestApi_GetCarts(t *testing.T) {
	var cart = domain.Cart{Description: "Test description"}
	service.CreateOrUpdate(&cart)
	req, err := http.NewRequest("GET", "/carts", nil)

	if err != nil {
		t.Fatal(err)
	}

	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
	var cartList[]domain.Cart
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

func getCartContent() domain.Cart {
	var items = map[string]domain.Item{
		"1":{
			Id:"1",
			Title:"Banana",
			Price:"2.50",
			Quantity: 1,
		},
		"2":{
			Id:"2",
			Title:"Apple",
			Price:"3.20",
			Quantity: 1,
		},
	}
	var cart = domain.Cart{
		Description: "Test cart",
		Items: items,
	}

	return cart
}


