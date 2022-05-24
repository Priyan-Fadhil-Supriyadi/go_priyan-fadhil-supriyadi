package controller

import (
	"github.com/Priyan-Fadhil-Supriyadi/mini-project/model"
	"github.com/stretchr/testify/mock"
)

type MockCarSvc struct {
	mock.Mock
}

func (m *MockCarSvc) CreateCarService(car model.Car) error {
	ret := m.Called()

	return ret.Error(0)
}
func (m *MockCarSvc) UpdateCarService(id int, car model.Car) error {
	return nil
}
func (m *MockCarSvc) GetAllCarService() []model.Car {
	return []model.Car{}
}
func (m *MockCarSvc) GetCarByID(id int) (model.Car, error) {
	return model.Car{}, nil
}
func (m *MockCarSvc) DeleteCarByID(id int) error {
	return nil
}
