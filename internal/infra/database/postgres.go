package database

import "github.com/brecabral/cattle-manager-mvp/internal/domain"

type AnimalDB struct{}

func (h *AnimalDB) Save(animal domain.Animal) (*domain.Animal, error) {
	return nil, nil
}

func (h *AnimalDB) FindByID() (*domain.Animal, error) {
	return nil, nil
}

func (h *AnimalDB) ListAll() ([]*domain.Animal, error) {
	return nil, nil
}

func (h *AnimalDB) Update(animal domain.Animal) error {
	return nil
}

func (h *AnimalDB) Delete() error {
	return nil
}
