package controller

import (
	"encoding/json"
	"net/http"
	"shopping-cart-service/src/main/domain"
)
type Api struct {
	ApiRouter http.Handler
}

func (a *Api) Router() http.Handler {
	return a.ApiRouter
}

func (a *Api) GetCarts(w http.ResponseWriter, r *http.Request) {
	var cart = domain.Cart{
		ID:    "01",
		Category:  "Category",
	}

	var content = []domain.Cart{cart}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(content)
}