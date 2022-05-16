package service

import (
	"errors"
	"testing"

	"github.com/Priyan-Fadhil-Supriyadi/mini-project/config"
	"github.com/Priyan-Fadhil-Supriyadi/mini-project/model"
	"github.com/stretchr/testify/assert"
)

func TestUpdateCar(t *testing.T) {
	testTable := []struct {
		name      string
		update    func(id int, car model.Car) error
		noError   bool
		token, id int
	}{
		{
			name: "success",
			update: func(id int, car model.Car) error {
				return nil
			},
			noError: true,
			token:   1,
			id:      1,
		},
		{
			name: "error token != id",
			update: func(id int, car model.Car) error {
				return nil
			},
			noError: false,
			token:   1,
			id:      2,
		},
		{
			name: "error internal",
			update: func(id int, car model.Car) error {
				return errors.New("error")
			},
			noError: false,
			token:   1,
			id:      1,
		},
	}
	repo := repoMockCar{}

	for _, v := range testTable {
		t.Run(v.name, func(t *testing.T) {
			repo.update = v.update

			svc := NewServiceCar(&repo, config.Config{})
			err := svc.UpdateCarService(v.id, v.token, model.Car{})
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}

func TestCreateCar(t *testing.T) {
	testTable := []struct {
		name    string
		create  func(car model.Car) error
		noError bool
	}{
		{
			name: "success",
			create: func(car model.Car) error {
				return nil
			},
			noError: true,
		},
		{
			name: "error internal",
			create: func(car model.Car) error {
				return errors.New("error")
			},
			noError: false,
		},
	}
	repo := repoMockCar{}

	for _, v := range testTable {
		t.Run(v.name, func(t *testing.T) {
			repo.create = v.create

			svc := NewServiceCar(&repo, config.Config{})
			err := svc.CreateCarService(model.Car{})
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}

func TestDeleteCar(t *testing.T) {
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
	repo := repoMockCar{}

	for _, v := range testTable {
		t.Run(v.name, func(t *testing.T) {
			repo.delete = v.delete

			svc := NewServiceCar(&repo, config.Config{})
			err := svc.DeleteCarByID(v.id)
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}

// func TestCreateCar(t *testing.T) {
// 	repo := repoMockCar2{}

// 	carController := svcCar{
// 		c:    config.Config{},
// 		repo: &repo,
// 	}

// 	car := model.Car{
// 		ID:    1,
// 		Type:  "Avanza",
// 		Brand: "Toyota",
// 		Price: "Rp 250.000",
// 	}

// 	repo.On("CreateTodoDetail", &car).Return(car)
// 	result := carController.CreateCarService(car)
// 	err := carController.CreateCarService(car)
// 	assert.Nil(t, err)
// 	assert.NotNil(t, result)
// }
