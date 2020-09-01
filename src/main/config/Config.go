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
	a.ApiRouter = r
	return a
}


