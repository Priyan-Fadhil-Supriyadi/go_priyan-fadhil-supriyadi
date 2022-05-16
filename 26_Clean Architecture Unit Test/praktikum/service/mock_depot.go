package service

import "github.com/Priyan-Fadhil-Supriyadi/mini-project/model"

type repoMockDepot struct {
	f      func(id int, depot model.Depot) error
	create func(depot model.Depot) error
	delete func(id int) error
}

func (r *repoMockDepot) CreateDepots(depot model.Depot) error {
	return r.create(depot)
}
func (r *repoMockDepot) GetAllDepots() []model.Depot {
	panic("implemente")
}
func (r *repoMockDepot) GetOneDepotByID(id int) (depot model.Depot, err error) {
	panic("implemente")
}
func (r *repoMockDepot) UpdateOneDepotByID(id int, depot model.Depot) error {
	return r.f(id, depot)
}
func (r *repoMockDepot) DeleteDepotByID(id int) error {
	return r.delete(id)
}
