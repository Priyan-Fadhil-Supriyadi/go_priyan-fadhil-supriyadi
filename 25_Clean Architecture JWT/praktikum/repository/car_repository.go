package repository

import (
	"fmt"

	"github.com/Priyan-Fadhil-Supriyadi/mini-project/domain"
	"github.com/Priyan-Fadhil-Supriyadi/mini-project/model"
	"gorm.io/gorm"
)

type carRepositoryMysqlLayer struct {
	DB *gorm.DB
}

func (r *carRepositoryMysqlLayer) CreateCars(car model.Car) error {
	res := r.DB.Create(&car)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func (r *carRepositoryMysqlLayer) GetAllCars() []model.Car {
	cars := []model.Car{}
	r.DB.Find(&cars)

	return cars
}

func (r *carRepositoryMysqlLayer) GetOneByIDCars(id int) (car model.Car, err error) {
	res := r.DB.Where("id = ?", id).Find(&car)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *carRepositoryMysqlLayer) GetOneByEmailCars(email string) (car model.Car, err error) {
	res := r.DB.Where("email = ?", email).Find(&car)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *carRepositoryMysqlLayer) UpdateOneByIDCars(id int, car model.Car) error {
	res := r.DB.Where("id = ?", id).UpdateColumns(&car)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update")
	}

	return nil
}

func (r *carRepositoryMysqlLayer) DeleteByIDCars(id int) error {
	res := r.DB.Delete(&model.Car{
		ID: id,
	})

	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete")
	}

	return nil
}

func CarMysqlRepository(db *gorm.DB) domain.AdapterCarRepository {
	return &carRepositoryMysqlLayer{
		DB: db,
	}
}
