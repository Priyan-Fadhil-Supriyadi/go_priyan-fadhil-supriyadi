package controller

import (
	"github.com/Priyan-Fadhil-Supriyadi/mini-project/model"
	"github.com/stretchr/testify/mock"
)

type MockRentSvc struct {
	mock.Mock
}

func (m *MockRentSvc) CreateRentService(rent model.Rent) error {
	ret := m.Called()

	return ret.Error(0)
}
func (m *MockRentSvc) UpdateRentService(id, idToken int, rent model.Rent) error {
	return nil
}
func (m *MockRentSvc) GetAllRentService() []model.Rent {
	return []model.Rent{}
}
func (m *MockRentSvc) GetRentByID(id int) (model.Rent, error) {
	return model.Rent{}, nil
}
func (m *MockRentSvc) DeleteRentByID(id int) error {
	return nil
}
