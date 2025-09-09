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

type User struct {
	ID           string `json:"id"`
	FullName     string `json:"fullName"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	IsShopeOwner string `json:"isShopeOwner"`
}

var UserList []User
