package global_product

type Fruits struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Color       string  `json:"color"`
	Image       string  `json:"image"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

var FruitsList []Fruits