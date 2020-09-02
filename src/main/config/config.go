package config

import (
	"github.com/gorilla/mux"
	"net/http"
	"shopping-cart-service/src/main/controller"
)

type Config interface {
	Router() http.Handler
}

func New() Config {
	a := &controller.Api{}
	r := mux.NewRouter()
	r.HandleFunc("/carts", a.GetCarts).Methods(http.MethodGet)
	r.HandleFunc("/carts", a.CreateCart).Methods(http.MethodPost)
	r.HandleFunc("/carts/{id}/items", a.AddItem).Methods(http.MethodPut)
	r.HandleFunc("/articles", a.GetArticles).Methods(http.MethodGet)
	a.ApiRouter = r
	return a
}


