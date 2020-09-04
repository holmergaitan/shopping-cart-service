package domain

type Item struct {
	Id    		string  `json:"id"`
	Title 		string  `json:"title"`
	Price 		string  `json:"price"`
	Quantity 	int 	`json:"quantity"`
}

func (i *Item) Increment()  {
	i.Quantity++
}
