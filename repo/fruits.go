package repo

import (
	"database/sql"
	"fruits-api/domain"
	"fruits-api/fruits"

	"github.com/jmoiron/sqlx"
)

type FruitsRepo interface {
	fruits.FruitsRepo
}

type fruitsRepo struct {
	db *sqlx.DB
}

// constructor function
func NewFruitsRepo(db *sqlx.DB) FruitsRepo {
	return &fruitsRepo{db: db}
}

// Create fruit
func (r *fruitsRepo) Create(f domain.Fruits) (*domain.Fruits, error) {
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
func (r *fruitsRepo) Get(fruitId int) (*domain.Fruits, error) {
	var f domain.Fruits
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
func (r *fruitsRepo) List(page, limit int64) ([]*domain.Fruits, error) {
	var fruits []*domain.Fruits

	offset := (page - 1) * limit

	query := `
		SELECT id, name, color, image, quantity, price, description
		FROM fruits
		ORDER BY id
		LIMIT $1 OFFSET $2
	`
	err := r.db.Select(&fruits, query, limit, offset)
	if err != nil {
		return nil, err
	}
	return fruits, nil
}

// Update fruit
func (r *fruitsRepo) Update(f domain.Fruits) (*domain.Fruits, error) {
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