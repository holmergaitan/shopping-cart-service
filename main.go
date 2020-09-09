package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net/http"
	"shopping-cart-service/src/main/shoppingcart"
)

func main() {
	dsn := "root:test@123@tcp(127.0.0.1:3306)/shopping_cart?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	_ = db.AutoMigrate(&shoppingcart.Item{})
	_ = db.AutoMigrate(&shoppingcart.Order{})
	_ = db.AutoMigrate(&shoppingcart.Cart{})
	var cartDao = &shoppingcart.CartDbRepository{Database: db}
	var itemDetailDao = &shoppingcart.OrderDbRepository{Database: db}
	var cache = &shoppingcart.ArticleDb{Database: db}
	cache.LoadAllItems()

	var mapper = shoppingcart.Mapper{}
	var service = shoppingcart.Service{
		CartDao: cartDao,
		OrderDao:itemDetailDao,
		ItemsCache: cache,
	}

	s := shoppingcart.New(service, mapper)
	log.Fatal(http.ListenAndServe(":8080", s.Router()))
}
