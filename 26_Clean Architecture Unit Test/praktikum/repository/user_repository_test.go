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

func TestGetAllUsers(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := UserMysqlRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "password", "nik"}).
			AddRow(1, "fadhil", "fadhil@gmail.com", "fadhil", "123456789").
			AddRow(2, "fadhil", "fadhil@gmail.com", "fadhil", "123456789").
			AddRow(3, "fadhil", "fadhil@gmail.com", "fadhil", "123456789"))

	res := repo.GetAllUsers()
	assert.Equal(t, res[0].Name, "fadhil")
	assert.Len(t, res, 3)
}

func TestDeleteUserByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := UserMysqlRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("DELETE")).
		WithArgs(1).
		WillReturnResult(driver.RowsAffected(1))
	fMock.ExpectCommit()

	err := repo.DeleteUserByID(1)
	assert.NoError(t, err)
	assert.True(t, true)
}

func TestUpdateUserByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := UserMysqlRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("UPDATE")).
		WithArgs("abc", 1).
		WillReturnResult(sqlmock.NewResult(0, 1))
	fMock.ExpectCommit()

	err := repo.UpdateOneUserByID(1, model.User{
		Name: "abc",
	})
	assert.NoError(t, err)
	assert.True(t, true)
}
