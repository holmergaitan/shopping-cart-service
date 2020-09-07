package shoppingcart

type Item struct {
	ID     string  `json:"id" gorm:"type:string;primary_key;"`
	Title  string  `json:"title"`
	Price  string  `json:"price"`
	Orders []Order `gorm:"foreignKey:ItemId"`
}

type Order struct{
	ID 			string		`json:"id" gorm:"type:string;primary_key;"`
	ItemId 		string		`json:"item"`
	CartId		string
	Quantity 	int			`json:"quantity"`
}

func (i *Order) Increment()  {
	i.Quantity++
}

type Cart struct {
	ID          string  `json:"id" gorm:"type:string;primary_key;"`
	Description string  `json:"description,omitempty" validate:"nil=false > empty=false" `
	Orders      []Order `json:"items" gorm:"foreignKey:CartId"`
}
