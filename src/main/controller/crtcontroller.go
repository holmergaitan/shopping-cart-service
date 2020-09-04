package controller

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"shopping-cart-service/src/main/domain"
	"shopping-cart-service/src/main/service"
)
type Api struct {
	ApiRouter http.Handler
}

func (a *Api) Router() http.Handler {
	return a.ApiRouter
}

func (a *Api) GetCarts(w http.ResponseWriter, r *http.Request) {
	var content = service.RetrieveCarts()
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(content)
}

func (a *Api) GetArticles(w http.ResponseWriter, r *http.Request){
	var content = service.LoadArticles()
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(content)
}

func (a *Api) CreateCart(w http.ResponseWriter, r *http.Request){
	cart := domain.Cart{Items: make(map[string]domain.Item)}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&cart); err != nil {
		buildErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()
	service.CreateOrUpdate(&cart)
	buildResponse(w, http.StatusCreated, cart)
}

func (a *Api) AddItem(w http.ResponseWriter, r *http.Request){
	service.LoadArticles()
	eventID := mux.Vars(r)["id"]
	item := domain.Item{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&item); err != nil {
		buildErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	var id, _ = uuid.Parse(eventID)
	var cart = service.GetCart(id)
	var articleInCache = service.ArticlesContent.Articles[item.Id]
	if val, ok := cart.Items[articleInCache.Id]; !ok{
		cart.Items[articleInCache.Id] = articleInCache
	}else{
		val.Increment()
	    cart.Items[articleInCache.Id] = val
	}
	service.CreateOrUpdate(cart)
	buildResponse(w, http.StatusCreated, cart)
}

func (a *Api) GetItems(w http.ResponseWriter, r *http.Request){
	cartId, _ := uuid.Parse(mux.Vars(r)["id"])
	var cart = service.GetCart(cartId)
	buildResponse(w, http.StatusCreated, cart.Items)
}

func (a *Api) DeleteItems(w http.ResponseWriter, r *http.Request){
	cartId, _ := uuid.Parse(mux.Vars(r)["id"])
	var cart = service.GetCart(cartId)
	cart.Items = make(map[string]domain.Item)
	service.CreateOrUpdate(cart)
	buildResponse(w, http.StatusCreated, cart)
}

func (a *Api) DeleteItem(w http.ResponseWriter, r *http.Request) {
	cartId, _ := uuid.Parse(mux.Vars(r)["id"])
	itemId, _ := mux.Vars(r)["itemId"]
	var cart = service.GetCart(cartId)
	delete(cart.Items, itemId)
	service.CreateOrUpdate(cart)
	buildResponse(w, http.StatusCreated, cart)
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