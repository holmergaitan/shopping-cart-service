package shoppingcart

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Config interface {
	Router() http.Handler
}

type Api struct {
	ApiRouter http.Handler
	Service Service
	Mapper Mapper
}

func New(service Service, mapper Mapper) Config {
	a := &Api{Service: service, Mapper: mapper}
	r := mux.NewRouter()
	r.HandleFunc("/carts", a.GetCarts).Methods(http.MethodGet)
	r.HandleFunc("/carts", a.SaveCart).Methods(http.MethodPost)
	r.HandleFunc("/carts/{id}/items", a.GetItems).Methods(http.MethodGet)
	r.HandleFunc("/carts/{id}/items", a.DeleteItems).Methods(http.MethodDelete)
	r.HandleFunc("/carts/{id}/items/{itemId}", a.DeleteItem).Methods(http.MethodDelete)
	r.HandleFunc("/carts/{id}/items", a.AddItem).Methods(http.MethodPut)
	r.HandleFunc("/articles", a.GetArticles).Methods(http.MethodGet)
	a.ApiRouter = r
	return a
}

func (a *Api) Router() http.Handler {
	return a.ApiRouter
}

func (a *Api) GetCarts(w http.ResponseWriter, r *http.Request) {
	var content = a.Service.GetCarts()
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(a.Mapper.ToCartDtoList(*content))
}

func (a *Api) GetArticles(w http.ResponseWriter, r *http.Request){
	var content = a.Service.ItemsCache.GetAll()
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(a.Mapper.ToItemDtoList(*content))
}

func (a *Api) SaveCart(w http.ResponseWriter, r *http.Request){
	cartDto := CartDto{Orders: make([]OrderDto, 0)}
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	if err := decoder.Decode(&cartDto); err != nil {
		buildErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	cart := a.Mapper.ToCarDomain(cartDto)
	_, err := a.Service.CreateCart(cart)
	if err != nil {
		buildErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	cartDto.ID = cart.ID
	buildResponse(w, http.StatusCreated, cartDto)
}

func (a *Api) AddItem(w http.ResponseWriter, r *http.Request){
	itemDto := ItemDto{}
	cartId := mux.Vars(r)["id"]
	var cart, carErr = a.Service.GetCart(cartId)
	if carErr != nil{
		buildErrorResponse(w, http.StatusNotFound,
			fmt.Sprintf("Cart with id %s not found", cartId))
		return
	}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	if err := decoder.Decode(&itemDto); err != nil {
		buildErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	var item, itemError = a.Service.GetItem(itemDto.ID)
	if itemError != nil{
		buildErrorResponse(w, http.StatusNotFound,
			fmt.Sprintf("Item with id %s not found", itemDto.ID))
		return
	}

	var order, _ = a.Service.GetOrderByCartAndItem(cartId, item.ID)
	if order.ID != "" {
		order.Increment()
		a.Service.OrderDao.UpdateOrder(order)
	}else {
		order = &Order{
			ItemId:     item.ID,
			CartId: 	cartId,
			Quantity: 	1,
		}
		_, orderErr := a.Service.CreateOrder(order)
		if orderErr != nil{
			buildErrorResponse(w, http.StatusInternalServerError, orderErr.Error())
			return
		}
	}

	var cartDto = a.Mapper.ToCartDto(*cart)
	buildResponse(w, http.StatusCreated, cartDto)
}

func (a *Api) GetItems(w http.ResponseWriter, r *http.Request){
	cartId, _ := mux.Vars(r)["id"]
	orders, err := a.Service.GetOrderByCartId(cartId)
	if err != nil {
		buildErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}
	orderDtoList := a.Mapper.ToOrderDtoList(*orders)
	buildResponse(w, http.StatusOK, orderDtoList)
}

func (a *Api) DeleteItems(w http.ResponseWriter, r *http.Request){
	cartId, _ := mux.Vars(r)["id"]
	err := a.Service.DeleteOrdersByCart(cartId)
	if err != nil {
		buildErrorResponse(w, http.StatusBadRequest,
			fmt.Sprintf("Can not delete items for cart with id %s", cartId))
		return
	}

	cart, cartErr := a.Service.GetCart(cartId)
	if cartErr != nil {
		buildErrorResponse(w, http.StatusBadRequest,
			fmt.Sprintf("Cart with id %s not found", cartId))
		return
	}

	cartDto := a.Mapper.ToCartDto(*cart)
	buildResponse(w, http.StatusOK, cartDto)

}

func (a *Api) DeleteItem(w http.ResponseWriter, r *http.Request) {
	cartId, _ := mux.Vars(r)["id"]
	itemId, _ := mux.Vars(r)["itemId"]
	order, err := a.Service.GetOrderByCartAndItem(cartId, itemId)
	if err != nil{
		buildErrorResponse(w, http.StatusBadRequest,
			fmt.Sprintf("Order not found"))
		return
	}

	if order.Quantity == 1{
		a.Service.DeleteOrder(order)
	}else{
		order.Remove()
		a.Service.UpdateOrder(order)
	}
	cart, _ := a.Service.GetCart(cartId)
	cartDto := a.Mapper.ToCartDto(*cart)
	buildResponse(w, http.StatusOK, cartDto)
}

func buildResponse(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

func buildErrorResponse(w http.ResponseWriter, code int, message string) {
	buildResponse(w, code, map[string]string{"error": message})
}

func toCartDto(){

}