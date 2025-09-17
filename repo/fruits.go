package repo

type Fruits struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Color       string  `json:"color"`
	Image       string  `json:"image"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

type FruitsRepo interface {
	Create(p Fruits) (*Fruits, error)
	Get(fruitId int) (*Fruits, error)
	List() ([]*Fruits, error)
	Delete(fruitId int) error
	Update(p Fruits) (*Fruits, error)
}

type fruitsRepo struct {
	FruitsList []*Fruits
}

func NewFruitsRepo() FruitsRepo {
	repo := &fruitsRepo{}
	generateInitFruits(repo)
	return repo
}

// Create
func (r *fruitsRepo) Create(f Fruits) (*Fruits, error) {
	f.ID = len(r.FruitsList) + 1
	r.FruitsList = append(r.FruitsList, &f)
	return &f, nil
}

// Get
func (r *fruitsRepo) Get(fruitId int) (*Fruits, error) {
	for _, fruit := range r.FruitsList {
		if fruit.ID == fruitId {
			return fruit, nil
		}
	}
	return nil, nil
}

// List
func (r *fruitsRepo) List() ([]*Fruits, error) {
	return r.FruitsList, nil
}

// Delete
func (r *fruitsRepo) Delete(fruitId int) error {
	var tempList []*Fruits
	for _, fruit := range r.FruitsList {
		if fruit.ID != fruitId {
			tempList = append(tempList, fruit)
		}
	}
	r.FruitsList = tempList
	return nil
}

// Update
func (r *fruitsRepo) Update(f Fruits) (*Fruits, error) {
	for index, fruit := range r.FruitsList {
		if fruit.ID == f.ID {
			r.FruitsList[index] = &f
			return &f, nil
		}
	}
	return nil, nil
}

// Initial Fruits
func generateInitFruits(r *fruitsRepo) {
	r.FruitsList = append(r.FruitsList, &Fruits{ID: 1, Name: "Apple", Color: "Red", Image: "https://images.unsplash.com/photo-1568702846914-96b305d2aaeb?w=500", Quantity: 50, Price: 10.99, Description: "Crisp and juicy red apples, perfect for snacking."})
	r.FruitsList = append(r.FruitsList, &Fruits{ID: 2, Name: "Banana", Color: "Yellow", Image: "https://images.unsplash.com/photo-1528825871115-3581a5387919?w=500", Quantity: 60, Price: 9.99, Description: "Sweet and creamy bananas, rich in potassium."})
	r.FruitsList = append(r.FruitsList, &Fruits{ID: 3, Name: "Blueberry", Color: "Blue", Image: "https://images.unsplash.com/photo-1498557850523-fd3d118b962e?w=500", Quantity: 70, Price: 8.99, Description: "Tiny antioxidant powerhouses with sweet-tart flavor."})
	r.FruitsList = append(r.FruitsList, &Fruits{ID: 4, Name: "Orange", Color: "Orange", Image: "https://images.unsplash.com/photo-1547514701-42782101795e?w=500", Quantity: 80, Price: 6.99, Description: "Juicy citrus fruits packed with vitamin C."})
	r.FruitsList = append(r.FruitsList, &Fruits{ID: 5, Name: "Kiwi", Color: "Green", Image: "https://images.unsplash.com/photo-1550258987-190a2d41a8ba?w=500", Quantity: 90, Price: 7.99, Description: "Fuzzy brown exterior with vibrant green flesh inside."})
	r.FruitsList = append(r.FruitsList, &Fruits{ID: 6, Name: "Strawberry", Color: "Red", Image: "https://images.unsplash.com/photo-1464965911861-746a04b4bca6?w=500", Quantity: 40, Price: 8.99, Description: "Sweet, heart-shaped berries perfect for desserts."})
}
