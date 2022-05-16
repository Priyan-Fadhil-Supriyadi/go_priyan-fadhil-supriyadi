package repository

import (
	"fmt"

	"github.com/Priyan-Fadhil-Supriyadi/mini-project/domain"
	"github.com/Priyan-Fadhil-Supriyadi/mini-project/model"
	"gorm.io/gorm"
)

type depotRepositoryMysqlLayer struct {
	DB *gorm.DB
}

func (r *depotRepositoryMysqlLayer) CreateDepots(depot model.Depot) error {
	res := r.DB.Create(&depot)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func (r *depotRepositoryMysqlLayer) GetAllDepots() []model.Depot {
	depots := []model.Depot{}
	r.DB.Find(&depots)

	return depots
}

func (r *depotRepositoryMysqlLayer) GetOneDepotByID(id int) (depot model.Depot, err error) {
	res := r.DB.Where("id = ?", id).Find(&depot)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *depotRepositoryMysqlLayer) UpdateOneDepotByID(id int, depot model.Depot) error {
	res := r.DB.Where("id = ?", id).UpdateColumns(&depot)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update")
	}

	return nil
}

func (r *depotRepositoryMysqlLayer) DeleteDepotByID(id int) error {
	res := r.DB.Delete(&model.Depot{
		ID: id,
	})

	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete")
	}

	return nil
}

func DepotMysqlRepository(db *gorm.DB) domain.AdapterDepotRepository {
	return &depotRepositoryMysqlLayer{
		DB: db,
	}
}
