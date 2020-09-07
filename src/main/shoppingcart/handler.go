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
	a.Service.CreateCart(cart)
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

	var cart = a.Service.GetCart(cartId)
	var cartDto = a.Mapper.ToCartDto(*cart)
	buildResponse(w, http.StatusCreated, cartDto)
}

func (a *Api) GetItems(w http.ResponseWriter, r *http.Request){
	cartId, _ := mux.Vars(r)["id"]
	//var cart = GetCart(cartId)
	buildResponse(w, http.StatusCreated, cartId)
}

func (a *Api) DeleteItems(w http.ResponseWriter, r *http.Request){
	cartId, _ := mux.Vars(r)["id"]
	//var cart = GetCart(cartId)
	//cart.Orders = make([]Order, 0)
	//CreateCart(cart)
	buildResponse(w, http.StatusCreated, cartId)
}

func (a *Api) DeleteItem(w http.ResponseWriter, r *http.Request) {
	//cartId, _ := uuid.Parse(mux.Vars(r)["id"])
	//itemId, _ := mux.Vars(r)["itemId"]
	//var cart = GetCart(cartId)
	////delete(cart.Orders, itemId)
	//Update(cart)
	//buildResponse(w, http.StatusCreated, cart)
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