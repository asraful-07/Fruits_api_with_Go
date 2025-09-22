package repo

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type User struct {
	ID           int    `db:"id" json:"id"`
	FullName     string `db:"full_name" json:"fullName"`
	Email        string `db:"email" json:"email"`
	Password     string `db:"password" json:"password"`
	IsShopeOwner bool   `db:"is_shope_owner" json:"isShopeOwner"`
}

type UserRepo interface {
	Create(u User) (*User, error)
	Find(email, pass string) (*User, error)
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{db: db}
}

// Create user
func (r *userRepo) Create(user User) (*User, error) {
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
func (r *userRepo) Find(email, pass string) (*User, error) {
	var user User
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
