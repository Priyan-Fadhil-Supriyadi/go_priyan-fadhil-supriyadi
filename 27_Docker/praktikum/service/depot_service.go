package service

import (
	"github.com/Priyan-Fadhil-Supriyadi/mini-project/config"
	"github.com/Priyan-Fadhil-Supriyadi/mini-project/domain"
	"github.com/Priyan-Fadhil-Supriyadi/mini-project/model"
)

type svcDepot struct {
	c    config.Config
	repo domain.AdapterDepotRepository
}

func (s *svcDepot) CreateDepotService(depot model.Depot) error {
	return s.repo.CreateDepots(depot)
}

func (s *svcDepot) UpdateDepotService(id int, depot model.Depot) error {
	return s.repo.UpdateOneDepotByID(id, depot)
}

func (s *svcDepot) GetAllDepotService() []model.Depot {
	return s.repo.GetAllDepots()
}

func (s *svcDepot) GetDepotByID(id int) (model.Depot, error) {
	return s.repo.GetOneDepotByID(id)
}

func (s *svcDepot) DeleteDepotByID(id int) error {
	return s.repo.DeleteDepotByID(id)
}

func NewServiceDepot(repo domain.AdapterDepotRepository, c config.Config) domain.AdapterDepotService {
	return &svcDepot{
		repo: repo,
		c:    c,
	}
}
