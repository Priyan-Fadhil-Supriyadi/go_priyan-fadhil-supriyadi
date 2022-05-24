package repository

import (
	"database/sql/driver"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Priyan-Fadhil-Supriyadi/mini-project/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestGetAllCars(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := CarMysqlRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `cars`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "type", "brand", "price"}).
			AddRow(1, "avanza", "toyota", "Rp 250.000").
			AddRow(2, "avanza", "toyota", "Rp 250.000").
			AddRow(3, "avanza", "toyota", "Rp 250.000"))

	res := repo.GetAllCars()
	assert.Equal(t, res[0].Type, "avanza")
	assert.Len(t, res, 3)
}

func TestDeleteCarByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := CarMysqlRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("DELETE")).
		WithArgs(1).
		WillReturnResult(driver.RowsAffected(1))
	fMock.ExpectCommit()

	err := repo.DeleteCarByID(1)
	assert.NoError(t, err)
	assert.True(t, true)
}

func TestUpdateCarByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := CarMysqlRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("UPDATE")).
		WithArgs("avanza", 1).
		WillReturnResult(sqlmock.NewResult(0, 1))
	fMock.ExpectCommit()

	err := repo.UpdateOneCarByID(1, model.Car{
		Type: "avanza",
	})
	assert.NoError(t, err)
	assert.True(t, true)
}
