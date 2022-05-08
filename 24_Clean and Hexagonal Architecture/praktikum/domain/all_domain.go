package domain

import (
	"github.com/labstack/echo/v4"
	"praktikum/model"
)

type AdapterRepository interface {
	CreateUsers(user model.User) error
	GetAll() []model.User
	GetOneByID(id int) (user model.User, err error)
	UpdateOneByID(id int, user model.User) error
	DeleteByID(id int) error
}

type AdapterController interface {
	CreateUserController(c echo.Context) error
	UpdateUserController(c echo.Context) error
	DeleteUserController(c echo.Context) error
	GetUserController(c echo.Context) error
	GetUsersController(c echo.Context) error
}
