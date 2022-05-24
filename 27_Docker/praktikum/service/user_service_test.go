package service

import (
	"errors"
	"testing"

	"github.com/Priyan-Fadhil-Supriyadi/mini-project/config"
	"github.com/Priyan-Fadhil-Supriyadi/mini-project/model"
	"github.com/stretchr/testify/assert"
)

func TestUpdateUser(t *testing.T) {
	testTable := []struct {
		name      string
		f         func(id int, user model.User) error
		noError   bool
		token, id int
	}{
		{
			name: "success",
			f: func(id int, user model.User) error {
				return nil
			},
			noError: true,
			token:   1,
			id:      1,
		},
		{
			name: "error token != id",
			f: func(id int, user model.User) error {
				return nil
			},
			noError: false,
			token:   1,
			id:      2,
		},
		{
			name: "error internal",
			f: func(id int, user model.User) error {
				return errors.New("error")
			},
			noError: false,
			token:   1,
			id:      1,
		},
	}
	repo := repoMockUser{}

	for _, v := range testTable {
		t.Run(v.name, func(t *testing.T) {
			repo.f = v.f

			svc := NewServiceUser(&repo, config.Config{})
			err := svc.UpdateUserService(v.id, v.token, model.User{})
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}

func TestCreateUser(t *testing.T) {
	testTable := []struct {
		name    string
		create  func(user model.User) error
		noError bool
	}{
		{
			name: "success",
			create: func(user model.User) error {
				return nil
			},
			noError: true,
		},
		{
			name: "error internal",
			create: func(user model.User) error {
				return errors.New("error")
			},
			noError: false,
		},
	}
	repo := repoMockUser{}

	for _, v := range testTable {
		t.Run(v.name, func(t *testing.T) {
			repo.create = v.create

			svc := NewServiceUser(&repo, config.Config{})
			err := svc.CreateUserService(model.User{})
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {
	testTable := []struct {
		name    string
		delete  func(id int) error
		noError bool
		id      int
	}{
		{
			name: "success",
			delete: func(id int) error {
				return nil
			},
			noError: true,
			id:      1,
		},

		{
			name: "error internal",
			delete: func(id int) error {
				return errors.New("error")
			},
			noError: false,
			id:      1,
		},
	}
	repo := repoMockUser{}

	for _, v := range testTable {
		t.Run(v.name, func(t *testing.T) {
			repo.delete = v.delete

			svc := NewServiceUser(&repo, config.Config{})
			err := svc.DeleteUserByID(v.id)
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}
