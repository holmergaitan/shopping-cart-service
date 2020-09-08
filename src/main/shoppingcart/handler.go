package shoppingcart

import (
	"encoding/json"
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
	r.HandleFunc("/carts", a.CreateCart).Methods(http.MethodPost)
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

func (a *Api) CreateCart(w http.ResponseWriter, r *http.Request){
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
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	if err := decoder.Decode(&itemDto); err != nil {
		buildErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	var item = a.Mapper.ToItemDomain(itemDto)
	var order = a.Service.GetOrderByCartAndItem(cartId, item.ID)
	if order.ID != "" {
		order.Increment()
		a.Service.OrderDao.Update(order)
	}else {
		order = &Order{
			ItemId:     item.ID,
			CartId: 	cartId,
			Quantity: 	1,
		}
		a.Service.CreateOrder(order)
	}

	var cart, _ = a.Service.GetCart(cartId)
	var cartDto = a.Mapper.ToCartDto(*cart)
	buildResponse(w, http.StatusCreated, cartDto)
}

func (a *Api) GetItems(w http.ResponseWriter, r *http.Request){
	cartId, _ := mux.Vars(r)["id"]
	orders := a.Service.GetOrderByCartId(cartId)
	orderDtoList := a.Mapper.ToOrderDtoList(*orders)
	buildResponse(w, http.StatusCreated, orderDtoList)
}

func (a *Api) DeleteItems(w http.ResponseWriter, r *http.Request){
	cartId, _ := mux.Vars(r)["id"]
	deleted := a.Service.DeleteOrdersByCart(cartId)
	cart, _ := a.Service.GetCart(cartId)
	cartDto := a.Mapper.ToCartDto(*cart)
	if deleted {
		buildResponse(w, http.StatusOK, cartDto)
	}else{
		buildErrorResponse(w, http.StatusBadRequest, "Can not delete items")
	}
}

func (a *Api) DeleteItem(w http.ResponseWriter, r *http.Request) {
	cartId, _ := mux.Vars(r)["id"]
	itemId, _ := mux.Vars(r)["itemId"]
	order := a.Service.GetOrderByCartAndItem(cartId, itemId)
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

func validateCartContent(){

}