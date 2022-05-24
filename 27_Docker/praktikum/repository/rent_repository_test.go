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

func TestGetAllRents(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := RentMysqlRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `rents`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "start", "deadline", "total"}).
			AddRow(1, "2 Mei 2022", "3 Mei 2022", 250.000).
			AddRow(2, "2 Mei 2022", "3 Mei 2022", 250.000).
			AddRow(3, "2 Mei 2022", "3 Mei 2022", 250.000))

	res := repo.GetAllRents()
	assert.Equal(t, res[0].Start, "2 Mei 2022")
	assert.Len(t, res, 3)
}

func TestDeleteRentByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := RentMysqlRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("DELETE")).
		WithArgs(1).
		WillReturnResult(driver.RowsAffected(1))
	fMock.ExpectCommit()

	err := repo.DeleteRentByID(1)
	assert.NoError(t, err)
	assert.True(t, true)
}

func TestUpdateRentByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := RentMysqlRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("UPDATE")).
		WithArgs("2 Mei 2022", 1).
		WillReturnResult(sqlmock.NewResult(0, 1))
	fMock.ExpectCommit()

	err := repo.UpdateOneRentByID(1, model.Rent{
		Start: "2 Mei 2022",
	})
	assert.NoError(t, err)
	assert.True(t, true)
}
