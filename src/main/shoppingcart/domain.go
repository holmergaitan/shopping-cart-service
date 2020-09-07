package shoppingcart

type Item struct {
	ID     string  `gorm:"type:string;primary_key;"`
	Title  string
	Price  string
	Orders []Order `gorm:"foreignKey:ItemId"`
}

type Order struct{
	ID 			string		`gorm:"type:string;primary_key;"`
	ItemId 		string
	CartId		string
	Quantity 	int
}

func (i *Order) Increment()  {
	i.Quantity++
}

type Cart struct {
	ID          string  `gorm:"type:string;primary_key;"`
	Description string
	Orders      []Order
}
