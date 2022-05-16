package controller

import (
	"github.com/Priyan-Fadhil-Supriyadi/mini-project/model"
	"github.com/stretchr/testify/mock"
)

type MockUserSvc struct {
	mock.Mock
}

func (m *MockUserSvc) CreateUserService(user model.User) error {
	ret := m.Called()

	return ret.Error(0)
}
func (m *MockUserSvc) UpdateUserService(id, idToken int, user model.User) error {
	return nil
}
func (m *MockUserSvc) GetAllUsersService() []model.User {
	return []model.User{}
}
func (m *MockUserSvc) GetUserByID(id int) (model.User, error) {
	return model.User{}, nil
}
func (m *MockUserSvc) LoginUser(email, password string) (string, int) {
	return "success", 200
}
func (m *MockUserSvc) DeleteUserByID(id int) error {
	return nil
}
