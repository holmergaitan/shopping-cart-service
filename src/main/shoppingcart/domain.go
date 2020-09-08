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
	Item 		Item        `gorm:"foreignKey:ItemId"`
	Quantity 	int
}

func (i *Order) Increment()  {
	i.Quantity++
}

func (i *Order) Remove()  {
	i.Quantity--
}

type Cart struct {
	ID          string  `gorm:"type:string;primary_key;"`
	Description string
	Orders      []Order `gorm:"foreignKey:CartId"`
}
