package service

import "github.com/Priyan-Fadhil-Supriyadi/mini-project/model"

type repoMockRent struct {
	f      func(id int, rent model.Rent) error
	create func(rent model.Rent) error
	delete func(id int) error
}

func (r *repoMockRent) CreateRents(rent model.Rent) error {
	return r.create(rent)
}
func (r *repoMockRent) GetAllRents() []model.Rent {
	panic("implemente")
}
func (r *repoMockRent) GetOneRentByID(id int) (rent model.Rent, err error) {
	panic("implemente")
}
func (r *repoMockRent) UpdateOneRentByID(id int, rent model.Rent) error {
	return r.f(id, rent)
}
func (r *repoMockRent) DeleteRentByID(id int) error {
	return r.delete(id)
}
