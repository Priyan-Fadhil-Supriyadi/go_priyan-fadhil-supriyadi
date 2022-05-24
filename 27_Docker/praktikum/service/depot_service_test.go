package service

import (
	"errors"
	"testing"

	"github.com/Priyan-Fadhil-Supriyadi/mini-project/config"
	"github.com/Priyan-Fadhil-Supriyadi/mini-project/model"
	"github.com/stretchr/testify/assert"
)

func TestUpdateDepot(t *testing.T) {
	testTable := []struct {
		name      string
		f         func(id int, depot model.Depot) error
		noError   bool
		token, id int
	}{
		{
			name: "success",
			f: func(id int, depot model.Depot) error {
				return nil
			},
			noError: true,
			token:   1,
			id:      1,
		},
		{
			name: "error token != id",
			f: func(id int, depot model.Depot) error {
				return nil
			},
			noError: false,
			token:   1,
			id:      2,
		},
		{
			name: "error internal",
			f: func(id int, depot model.Depot) error {
				return errors.New("error")
			},
			noError: false,
			token:   1,
			id:      1,
		},
	}
	repo := repoMockDepot{}

	for _, v := range testTable {
		t.Run(v.name, func(t *testing.T) {
			repo.f = v.f

			svc := NewServiceDepot(&repo, config.Config{})
			err := svc.UpdateDepotService(v.id, model.Depot{})
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}

func TestCreateDepot(t *testing.T) {
	testTable := []struct {
		name    string
		create  func(depot model.Depot) error
		noError bool
	}{
		{
			name: "success",
			create: func(depot model.Depot) error {
				return nil
			},
			noError: true,
		},
		{
			name: "error internal",
			create: func(depot model.Depot) error {
				return errors.New("error")
			},
			noError: false,
		},
	}
	repo := repoMockDepot{}

	for _, v := range testTable {
		t.Run(v.name, func(t *testing.T) {
			repo.create = v.create

			svc := NewServiceDepot(&repo, config.Config{})
			err := svc.CreateDepotService(model.Depot{})
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}

func TestDeleteDepot(t *testing.T) {
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
	repo := repoMockDepot{}

	for _, v := range testTable {
		t.Run(v.name, func(t *testing.T) {
			repo.delete = v.delete

			svc := NewServiceDepot(&repo, config.Config{})
			err := svc.DeleteDepotByID(v.id)
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}
