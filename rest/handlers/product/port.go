package product

import "fruits-api/domain"

type Service interface {
	Create(fruit domain.Fruits) (*domain.Fruits, error)
	Get(id int) (*domain.Fruits, error)
	List() ([]*domain.Fruits, error)
	Update(domain.Fruits) (*domain.Fruits, error)
	Delete(id int) error
}