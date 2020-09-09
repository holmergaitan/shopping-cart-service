package shoppingcart

import (
	"encoding/json"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"net/http"
)

type CacheInterface interface {
	LoadAllItems()

	GetItemById(id string) (*Item, error)

	GetAllItems() *[]Item

	LoadItemsByCartId(cartId string) []Item
}

type ArticleDb struct {
	Database *gorm.DB
}

type ArticlesCache struct {
	Articles map[string]Item
	Loaded bool
}

func (db *ArticleDb) LoadAllItems() {
	var articleList = load()
	db.Database.Create(&articleList)
}

func (db *ArticleDb) GetItemById(id string) (*Item, error) {
	var item Item
	result := db.Database.First(&item, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &item, nil
}

func (db *ArticleDb) GetAllItems() *[]Item {
	var items = make([]Item, 0)
	db.Database.Find(&items)
	return &items
}

func (db *ArticleDb) LoadItemsByCartId(cartId string)[]Item {
	var items = make([]Item, 0)
	db.Database.
		Select("*").
		Joins("INNER JOIN orders ON orders.item_id = items.id").
		Where("orders.cart_id = ?", cartId).Find(&items)

	return items
}

func (cache *ArticlesCache) Load(){
	var articleList = load()
	if !cache.Loaded{
		for _, art := range articleList {
			cache.Articles[art.ID] = art
		}
	}
	cache.Loaded = true
}

func (cache *ArticlesCache) GetById(id string) *Item {
	return nil
}

func (cache *ArticlesCache) GetAll() []Item {
	var items = make([]Item, 0)
	return items
}

func load() []Item {
	var articles []Item
	resp, err := http.Get("http://challenge.getsandbox.com/articles")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	log.Print("API Response as String:\n" + bodyString)
	json.Unmarshal(bodyBytes, &articles)
	log.Printf("API Response as struct %+v\n", articles)

	return articles
}