package domain

import (
	"github.com/Priyan-Fadhil-Supriyadi/mini-project/model"
)

type AdapterUserRepository interface {
	CreateUsers(user model.User) error
	GetAllUsers() []model.User
	GetOneByIDUsers(id int) (user model.User, err error)
	GetOneByEmailUsers(email string) (user model.User, err error)
	UpdateOneByIDUsers(id int, user model.User) error
	DeleteByIDUsers(id int) error
}

type AdapterCarRepository interface {
	CreateCars(car model.Car) error
	GetAllCars() []model.Car
	GetOneByIDCars(id int) (car model.Car, err error)
	GetOneByEmailCars(email string) (car model.Car, err error)
	UpdateOneByIDCars(id int, car model.Car) error
	DeleteByIDCars(id int) error
}

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
	UpdateCarService(id, idToken int, car model.Car) error
	GetAllCarsService() []model.Car
	GetCarByID(id int) (model.Car, error)
	DeleteCarByID(id int) error
}
