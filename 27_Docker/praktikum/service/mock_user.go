package service

import "github.com/Priyan-Fadhil-Supriyadi/mini-project/model"

type repoMockUser struct {
	f      func(id int, user model.User) error
	create func(user model.User) error
	delete func(id int) error
}

func (r *repoMockUser) CreateUsers(user model.User) error {
	return r.create(user)
}
func (r *repoMockUser) GetAllUsers() []model.User {
	panic("implemente")
}
func (r *repoMockUser) GetOneUserByID(id int) (user model.User, err error) {
	panic("implemente")
}
func (r *repoMockUser) GetOneUserByEmail(email string) (user model.User, err error) {
	panic("implemente")
}
func (r *repoMockUser) UpdateOneUserByID(id int, user model.User) error {
	return r.f(id, user)
}
func (r *repoMockUser) DeleteUserByID(id int) error {
	return r.delete(id)
}
