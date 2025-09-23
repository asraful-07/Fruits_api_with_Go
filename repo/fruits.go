package repo

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type Fruits struct {
	ID          int     `db:"id" json:"id"`
	Name        string  `db:"name" json:"name"`
	Color       string  `db:"color" json:"color"`
	Image       string  `db:"image" json:"image"`
	Quantity    int     `db:"quantity" json:"quantity"`
	Price       float64 `db:"price" json:"price"`
	Description string  `db:"description" json:"description"`
}

type FruitsRepo interface {
	Create(f Fruits) (*Fruits, error)
	Get(fruitId int) (*Fruits, error)
	List() ([]*Fruits, error)
	Update(f Fruits) (*Fruits, error)
	Delete(fruitId int) error
}

type fruitsRepo struct {
	db *sqlx.DB
}

func NewFruitsRepo(db *sqlx.DB) FruitsRepo {
	return &fruitsRepo{db: db}
}

// Create fruit
func (r *fruitsRepo) Create(f Fruits) (*Fruits, error) {
	query := `
		INSERT INTO fruits (name, color, image, quantity, price, description)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`

	err := r.db.QueryRow(
		query,
		f.Name,
		f.Color,
		f.Image,
		f.Quantity,
		f.Price,
		f.Description,
	).Scan(&f.ID)

	if err != nil {
		return nil, err
	}

	return &f, nil
}

// Get fruit by id
func (r *fruitsRepo) Get(fruitId int) (*Fruits, error) {
	var f Fruits
	query := `
		SELECT id, name, color, image, quantity, price, description
		FROM fruits
		WHERE id = $1
	`
	err := r.db.Get(&f, query, fruitId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &f, nil
}

// List all fruits
func (r *fruitsRepo) List() ([]*Fruits, error) {
	var fruits []*Fruits
	query := `
		SELECT id, name, color, image, quantity, price, description
		FROM fruits
		ORDER BY id
	`
	err := r.db.Select(&fruits, query)
	if err != nil {
		return nil, err
	}
	return fruits, nil
}

// Update fruit
func (r *fruitsRepo) Update(f Fruits) (*Fruits, error) {
	query := `
		UPDATE fruits
		SET name = :name,
			color = :color,
			image = :image,
			quantity = :quantity,
			price = :price,
			description = :description
		WHERE id = :id
		RETURNING id
	`

	rows, err := r.db.NamedQuery(query, f)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(&f.ID); err != nil {
			return nil, err
		}
	}

	return &f, nil
}

// Delete fruit
func (r *fruitsRepo) Delete(fruitId int) error {
	query := `DELETE FROM fruits WHERE id = $1`
	_, err := r.db.Exec(query, fruitId)
	return err
}

 /*
// Initial Fruits
func generateInitFruits(r *fruitsRepo) {
	r.FruitsList = append(r.FruitsList, &Fruits{ID: 1, Name: "Apple", Color: "Red", Image: "https://images.unsplash.com/photo-1568702846914-96b305d2aaeb?w=500", Quantity: 50, Price: 10.99, Description: "Crisp and juicy red apples, perfect for snacking."})
	r.FruitsList = append(r.FruitsList, &Fruits{ID: 2, Name: "Banana", Color: "Yellow", Image: "https://images.unsplash.com/photo-1528825871115-3581a5387919?w=500", Quantity: 60, Price: 9.99, Description: "Sweet and creamy bananas, rich in potassium."})
	r.FruitsList = append(r.FruitsList, &Fruits{ID: 3, Name: "Blueberry", Color: "Blue", Image: "https://images.unsplash.com/photo-1498557850523-fd3d118b962e?w=500", Quantity: 70, Price: 8.99, Description: "Tiny antioxidant powerhouses with sweet-tart flavor."})
	r.FruitsList = append(r.FruitsList, &Fruits{ID: 4, Name: "Orange", Color: "Orange", Image: "https://images.unsplash.com/photo-1547514701-42782101795e?w=500", Quantity: 80, Price: 6.99, Description: "Juicy citrus fruits packed with vitamin C."})
	r.FruitsList = append(r.FruitsList, &Fruits{ID: 5, Name: "Kiwi", Color: "Green", Image: "https://images.unsplash.com/photo-1550258987-190a2d41a8ba?w=500", Quantity: 90, Price: 7.99, Description: "Fuzzy brown exterior with vibrant green flesh inside."})
	r.FruitsList = append(r.FruitsList, &Fruits{ID: 6, Name: "Strawberry", Color: "Red", Image: "https://images.unsplash.com/photo-1464965911861-746a04b4bca6?w=500", Quantity: 40, Price: 8.99, Description: "Sweet, heart-shaped berries perfect for desserts."})
}
*/