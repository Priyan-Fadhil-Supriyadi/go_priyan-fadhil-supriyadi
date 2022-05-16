package service

import (
	"fmt"

	"github.com/Priyan-Fadhil-Supriyadi/mini-project/config"
	"github.com/Priyan-Fadhil-Supriyadi/mini-project/domain"
	"github.com/Priyan-Fadhil-Supriyadi/mini-project/model"
)

type svcCar struct {
	c    config.Config
	repo domain.AdapterCarRepository
}

func (s *svcCar) CreateCarService(car model.Car) error {
	return s.repo.CreateCars(car)
}

func (s *svcCar) UpdateCarService(id, idToken int, car model.Car) error {
	fmt.Println(id, idToken)
	if id != idToken {
		return fmt.Errorf("error")
	}
	return s.repo.UpdateOneCarByID(id, car)
}

func (s *svcCar) GetAllCarService() []model.Car {
	return s.repo.GetAllCars()
}

func (s *svcCar) GetCarByID(id int) (model.Car, error) {
	return s.repo.GetOneCarByID(id)
}

func (s *svcCar) DeleteCarByID(id int) error {
	return s.repo.DeleteCarByID(id)
}

func NewServiceCar(repo domain.AdapterCarRepository, c config.Config) domain.AdapterCarService {
	return &svcCar{
		repo: repo,
		c:    c,
	}
}
