package fruits

import (
	"fruits-api/domain"
	"fruits-api/rest/handlers/product"
)

type Service interface {
	product.Service
}

type FruitsRepo interface {
	Create(f domain.Fruits) (*domain.Fruits, error)
	Get(fruitId int) (*domain.Fruits, error)
	List(page, limit int64) ([]*domain.Fruits, error)
	Update(f domain.Fruits) (*domain.Fruits, error)
	Delete(fruitId int) error
}
