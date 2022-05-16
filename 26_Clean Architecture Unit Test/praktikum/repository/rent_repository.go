package repository

import (
	"fmt"

	"github.com/Priyan-Fadhil-Supriyadi/mini-project/domain"
	"github.com/Priyan-Fadhil-Supriyadi/mini-project/model"
	"gorm.io/gorm"
)

type rentRepositoryMysqlLayer struct {
	DB *gorm.DB
}

func (r *rentRepositoryMysqlLayer) CreateRents(rent model.Rent) error {
	res := r.DB.Create(&rent)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func (r *rentRepositoryMysqlLayer) GetAllRents() []model.Rent {
	rents := []model.Rent{}
	r.DB.Find(&rents)

	return rents
}

func (r *rentRepositoryMysqlLayer) GetOneRentByID(id int) (rent model.Rent, err error) {
	res := r.DB.Where("id = ?", id).Find(&rent)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *rentRepositoryMysqlLayer) UpdateOneRentByID(id int, rent model.Rent) error {
	res := r.DB.Where("id = ?", id).UpdateColumns(&rent)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update")
	}

	return nil
}

func (r *rentRepositoryMysqlLayer) DeleteRentByID(id int) error {
	res := r.DB.Delete(&model.Rent{
		ID: id,
	})

	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete")
	}

	return nil
}

func RentMysqlRepository(db *gorm.DB) domain.AdapterRentRepository {
	return &rentRepositoryMysqlLayer{
		DB: db,
	}
}
