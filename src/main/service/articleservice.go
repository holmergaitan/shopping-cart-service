package service

import (
	"shopping-cart-service/src/main/client"
	"shopping-cart-service/src/main/domain"
)

var ArticlesContent = client.Cache

func LoadArticles() map[string]domain.Item {
	ArticlesContent.Load()
	return ArticlesContent.Articles
}