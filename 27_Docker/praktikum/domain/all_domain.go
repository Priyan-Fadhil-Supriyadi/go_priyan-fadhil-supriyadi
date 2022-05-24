package domain

import (
	"github.com/Priyan-Fadhil-Supriyadi/mini-project/model"
)

//repository
type AdapterUserRepository interface {
	CreateUsers(user model.User) error
	GetAllUsers() []model.User
	GetOneUserByID(id int) (user model.User, err error)
	GetOneUserByEmail(email string) (user model.User, err error)
	UpdateOneUserByID(id int, user model.User) error
	DeleteUserByID(id int) error
}

type AdapterCarRepository interface {
	CreateCars(car model.Car) error
	GetAllCars() []model.Car
	GetOneCarByID(id int) (car model.Car, err error)
	UpdateOneCarByID(id int, car model.Car) error
	DeleteCarByID(id int) error
}

type AdapterDepotRepository interface {
	CreateDepots(depot model.Depot) error
	GetAllDepots() []model.Depot
	GetOneDepotByID(id int) (depot model.Depot, err error)
	UpdateOneDepotByID(id int, depot model.Depot) error
	DeleteDepotByID(id int) error
}

type AdapterRentRepository interface {
	CreateRents(rent model.Rent) error
	GetAllRents() []model.Rent
	GetOneRentByID(id int) (rent model.Rent, err error)
	UpdateOneRentByID(id int, rent model.Rent) error
	DeleteRentByID(id int) error
}

// domain handler
type AdapterUserService interface {
	CreateUserService(user model.User) error
	UpdateUserService(id, idToken int, user model.User) error
	GetAllUsersService() []model.User
	GetUserByID(id int) (model.User, error)
	LoginUser(email, password string) (string, int)
	DeleteUserByID(id int) error
}

type AdapterCarService interface {
	CreateCarService(car model.Car) error
	UpdateCarService(id int, car model.Car) error
	GetAllCarService() []model.Car
	GetCarByID(id int) (model.Car, error)
	DeleteCarByID(id int) error
}

type AdapterDepotService interface {
	CreateDepotService(depot model.Depot) error
	UpdateDepotService(id int, depot model.Depot) error
	GetAllDepotService() []model.Depot
	GetDepotByID(id int) (model.Depot, error)
	DeleteDepotByID(id int) error
}

type AdapterRentService interface {
	CreateRentService(rent model.Rent) error
	UpdateRentService(id, idUser, idToken int, rent model.Rent) error
	GetAllRentService() []model.Rent
	GetRentByID(id int) (model.Rent, error)
	DeleteRentByID(id int) error
}
