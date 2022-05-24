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

// func TestCreateDepots(t *testing.T) {
// 	dbMock, fMock, _ := sqlmock.New()
// 	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
// 		Conn:                      dbMock,
// 		SkipInitializeWithVersion: true,
// 	},
// 	})
// 	repo := DepotMysqlRepository(db)
// 	defer dbMock.Close()

// 	fMock.ExpectQuery(regexp.QuoteMeta("INSERT INTO `cars`")).
// 		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "phone", "address", "carId"}).
// 			AddRow(1, "hari", "08561203123", "munjul jaya permai, purwakarta", 1))

// 	fMock.ExpectBegin()
// 	fMock.ExpectExec(regexp.QuoteMeta("CREATE")).
// 		WithArgs("hari").
// 		WillReturnResult(sqlmock.NewResult(0, 1))
// 	fMock.ExpectCommit()

// 	err := repo.CreateDepots(model.Depot{
// 		Name: "hari",
// 	})
// 	assert.NoError(t, err)
// 	assert.True(t, true)
// }

func TestGetAllDepots(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := DepotMysqlRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `depots`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "phone", "adress"}).
			AddRow(1, "hari", "08561203123", "munjul jaya permai, purwakarta").
			AddRow(2, "hari", "08561203123", "munjul jaya permai, purwakarta").
			AddRow(3, "hari", "08561203123", "munjul jaya permai, purwakarta"))

	res := repo.GetAllDepots()
	assert.Equal(t, res[0].Name, "hari")
	assert.Len(t, res, 3)
}

func TestDeleteDepotByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := DepotMysqlRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("DELETE")).
		WithArgs(1).
		WillReturnResult(driver.RowsAffected(1))
	fMock.ExpectCommit()

	err := repo.DeleteDepotByID(1)
	assert.NoError(t, err)
	assert.True(t, true)
}

func TestUpdateDepotByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := DepotMysqlRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("UPDATE")).
		WithArgs("hari", 1).
		WillReturnResult(sqlmock.NewResult(0, 1))
	fMock.ExpectCommit()

	err := repo.UpdateOneDepotByID(1, model.Depot{
		Name: "hari",
	})
	assert.NoError(t, err)
	assert.True(t, true)
}
