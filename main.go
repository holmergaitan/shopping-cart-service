package main

import (
	"log"
	"net/http"
	"shopping-cart-service/src/main/config"
)


func main() {
	s := config.New()
	log.Fatal(http.ListenAndServe(":8080", s.Router()))
}
