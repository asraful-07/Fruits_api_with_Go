package domain

// model or entity -> existence
type Fruits struct {
	ID          int     `db:"id" json:"id"`
	Name        string  `db:"name" json:"name"`
	Color       string  `db:"color" json:"color"`
	Image       string  `db:"image" json:"image"`
	Quantity    int     `db:"quantity" json:"quantity"`
	Price       float64 `db:"price" json:"price"`
	Description string  `db:"description" json:"description"`
}