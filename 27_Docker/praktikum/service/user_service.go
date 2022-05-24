package service

import (
	"fmt"
	"net/http"

	"github.com/Priyan-Fadhil-Supriyadi/mini-project/config"
	"github.com/Priyan-Fadhil-Supriyadi/mini-project/domain"
	"github.com/Priyan-Fadhil-Supriyadi/mini-project/helper"
	"github.com/Priyan-Fadhil-Supriyadi/mini-project/model"
)

type svcUser struct {
	c    config.Config
	repo domain.AdapterUserRepository
}

func (s *svcUser) CreateUserService(user model.User) error {
	return s.repo.CreateUsers(user)
}

func (s *svcUser) UpdateUserService(id, idToken int, user model.User) error {
	fmt.Println(id, idToken)
	if id != idToken {
		return fmt.Errorf("error")
	}
	return s.repo.UpdateOneUserByID(id, user)
}

func (s *svcUser) GetAllUsersService() []model.User {
	return s.repo.GetAllUsers()
}

func (s *svcUser) GetUserByID(id int) (model.User, error) {
	return s.repo.GetOneUserByID(id)
}

func (s *svcUser) LoginUser(email, password string) (string, int) {
	user, _ := s.repo.GetOneUserByEmail(email)
	fmt.Println(model.User{})

	if user.Password != password {
		return "", http.StatusUnauthorized
	}

	token, err := helper.CreateToken(user.ID, user.Email, s.c.JWT_KEY)
	if err != nil {
		return "", http.StatusInternalServerError
	}

	return token, http.StatusOK
}

func (s *svcUser) DeleteUserByID(id int) error {
	return s.repo.DeleteUserByID(id)
}

func NewServiceUser(repo domain.AdapterUserRepository, c config.Config) domain.AdapterUserService {
	return &svcUser{
		repo: repo,
		c:    c,
	}
}
