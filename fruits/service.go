package fruits

import "fruits-api/domain"

type service struct {
	fruitsRepo FruitsRepo
}

func NewService(fruitsRepo FruitsRepo) Service {
	return &service{
		fruitsRepo: fruitsRepo,
	}
}

func (svc service) Create(fruit domain.Fruits) (*domain.Fruits, error) {
	return svc.fruitsRepo.Create(fruit)
}

func (svc service) Get(id int) (*domain.Fruits, error) {
	return svc.fruitsRepo.Get(id)
}

func (svc service) List() ([]*domain.Fruits, error) {
	return  svc.fruitsRepo.List()
}

func (svc service) Update(fruit domain.Fruits) (*domain.Fruits, error) {
	return svc.fruitsRepo.Update(fruit)
}

func (svc service) Delete(id int) error {
	return  svc.fruitsRepo.Delete(id)
}