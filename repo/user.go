package repo

import (
	"database/sql"
	"fruits-api/domain"
	"fruits-api/user"

	"github.com/jmoiron/sqlx"
)

type UserRepo interface {
	user.UserRepo
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{db: db} 
}

// Create user
func (r *userRepo) Create(user domain.User) (*domain.User, error) {
	query := `
		INSERT INTO users (full_name, email, password, is_shope_owner)
		VALUES (:full_name, :email, :password, :is_shope_owner)
		RETURNING id
	`

	rows, err := r.db.NamedQuery(query, user)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(&user.ID); err != nil {
			return nil, err
		}
	}

	return &user, nil
}

// Find user by email & password
func (r *userRepo) Find(email, pass string) (*domain.User, error) {
	var user domain.User
	query := `
		SELECT id, full_name, email, password, is_shope_owner
		FROM users
		WHERE email = $1 AND password = $2
	`
	err := r.db.Get(&user, query, email, pass)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) List() ([]* domain.User, error)  {
	var users []*domain.User
	query := `
	SELECT id, full_name, email, password, is_shope_owner
	FROM users
	ORDER BY id 
	`

	err := r.db.Select(&users, query)

	if err != nil {
		return nil, err
	}
	return users, nil
}


func (r *userRepo) Delete(id int) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.Exec(query)
	return err
}