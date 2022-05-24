package service

import "github.com/Priyan-Fadhil-Supriyadi/mini-project/model"

type repoMockCar struct {
	update func(id int, car model.Car) error
	create func(car model.Car) error
	delete func(id int) error
}

func (r *repoMockCar) CreateCars(car model.Car) error {
	return r.create(car)
}
func (r *repoMockCar) GetAllCars() []model.Car {
	panic("implemente")
}
func (r *repoMockCar) GetOneCarByID(id int) (car model.Car, err error) {
	panic("implemente")
}
func (r *repoMockCar) UpdateOneCarByID(id int, car model.Car) error {
	return r.update(id, car)
}
func (r *repoMockCar) DeleteCarByID(id int) error {
	return r.delete(id)
}
