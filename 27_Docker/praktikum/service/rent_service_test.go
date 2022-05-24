package service

import (
	"errors"
	"testing"

	"github.com/Priyan-Fadhil-Supriyadi/mini-project/config"
	"github.com/Priyan-Fadhil-Supriyadi/mini-project/model"
	"github.com/stretchr/testify/assert"
)

func TestUpdateRent(t *testing.T) {
	testTable := []struct {
		name              string
		f                 func(id int, rent model.Rent) error
		noError           bool
		token, userId, id int
	}{
		{
			name: "success",
			f: func(id int, rent model.Rent) error {
				return nil
			},
			noError: true,
			token:   1,
			id:      1,
			userId:  1,
		},
		{
			name: "error token != id",
			f: func(id int, rent model.Rent) error {
				return nil
			},
			noError: false,
			token:   1,
			id:      2,
			userId:  3,
		},
		{
			name: "error internal",
			f: func(id int, rent model.Rent) error {
				return errors.New("error")
			},
			noError: false,
			token:   1,
			id:      1,
			userId:  1,
		},
	}
	repo := repoMockRent{}

	for _, v := range testTable {
		t.Run(v.name, func(t *testing.T) {
			repo.f = v.f

			svc := NewServiceRent(&repo, config.Config{})
			err := svc.UpdateRentService(v.id, v.token, v.userId, model.Rent{})
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}

func TestCreateRent(t *testing.T) {
	testTable := []struct {
		name    string
		create  func(rent model.Rent) error
		noError bool
	}{
		{
			name: "success",
			create: func(rent model.Rent) error {
				return nil
			},
			noError: true,
		},
		{
			name: "error internal",
			create: func(rent model.Rent) error {
				return errors.New("error")
			},
			noError: false,
		},
	}
	repo := repoMockRent{}

	for _, v := range testTable {
		t.Run(v.name, func(t *testing.T) {
			repo.create = v.create

			svc := NewServiceRent(&repo, config.Config{})
			err := svc.CreateRentService(model.Rent{})
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}

func TestDeleteRent(t *testing.T) {
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
	repo := repoMockRent{}

	for _, v := range testTable {
		t.Run(v.name, func(t *testing.T) {
			repo.delete = v.delete

			svc := NewServiceRent(&repo, config.Config{})
			err := svc.DeleteRentByID(v.id)
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}
