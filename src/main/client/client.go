package client

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"shopping-cart-service/src/main/domain"
)
type ArticlesCache struct {
	Articles map[string]domain.Item
	Loaded bool
}

type CacheInterface interface {
	Load()
}

var Cache = ArticlesCache{Articles: make(map[string]domain.Item)}

func (cache *ArticlesCache)Load(){
	var articleList = load()
	if !cache.Loaded{
		for _, art := range articleList {
			cache.Articles[art.Id] = art
		}
	}
	cache.Loaded = true
}

func load() []domain.Item {
	var articles []domain.Item
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