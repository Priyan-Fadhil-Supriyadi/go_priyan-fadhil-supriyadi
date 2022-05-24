package controller

import (
	"github.com/Priyan-Fadhil-Supriyadi/mini-project/model"
	"github.com/stretchr/testify/mock"
)

type MockDepotSvc struct {
	mock.Mock
}

func (m *MockDepotSvc) CreateDepotService(depot model.Depot) error {
	ret := m.Called()

	return ret.Error(0)
}
func (m *MockDepotSvc) UpdateDepotService(id int, depot model.Depot) error {
	return nil
}
func (m *MockDepotSvc) GetAllDepotService() []model.Depot {
	return []model.Depot{}
}
func (m *MockDepotSvc) GetDepotByID(id int) (model.Depot, error) {
	return model.Depot{}, nil
}
func (m *MockDepotSvc) DeleteDepotByID(id int) error {
	return nil
}
