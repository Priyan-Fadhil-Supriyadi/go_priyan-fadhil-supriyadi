package service

import (
	"fmt"

	"github.com/Priyan-Fadhil-Supriyadi/mini-project/config"
	"github.com/Priyan-Fadhil-Supriyadi/mini-project/domain"
	"github.com/Priyan-Fadhil-Supriyadi/mini-project/model"
)

type svcRent struct {
	c    config.Config
	repo domain.AdapterRentRepository
}

func (s *svcRent) CreateRentService(rent model.Rent) error {
	return s.repo.CreateRents(rent)
}

func (s *svcRent) UpdateRentService(id, idToken int, rent model.Rent) error {
	fmt.Println(id, idToken)
	if id != idToken {
		return fmt.Errorf("error")
	}
	return s.repo.UpdateOneRentByID(id, rent)
}

func (s *svcRent) GetAllRentService() []model.Rent {
	return s.repo.GetAllRents()
}

func (s *svcRent) GetRentByID(id int) (model.Rent, error) {
	return s.repo.GetOneRentByID(id)
}

func (s *svcRent) DeleteRentByID(id int) error {
	return s.repo.DeleteRentByID(id)
}

func NewServiceRent(repo domain.AdapterRentRepository, c config.Config) domain.AdapterRentService {
	return &svcRent{
		repo: repo,
		c:    c,
	}
}
