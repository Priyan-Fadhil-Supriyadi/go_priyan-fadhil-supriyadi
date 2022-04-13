package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"praktikum/database"
)

func TestHTTPBasic(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	A(w, r)

	assert.Equal(t, 500, w.Result().StatusCode)
}

func TestEchoGetAllUsers(t *testing.T) {
	db_mock, mk, _ := sqlmock.New()

	database.DB, _ = gorm.Open(mysql.New(
		mysql.Config{
			Conn:                      db_mock,
			SkipInitializeWithVersion: true,
		},
	))

	mk.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users`")).WillReturnRows(
		sqlmock.NewRows([]string{"id", "name", "email", "password"}).
			AddRow(1, "fadhil", "fadhil@gmail.com", "fadhil"),
	)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()

	e := echo.New()
	ct := e.NewContext(req, resp)

	getUsersController(ct)

	res := make(map[string]interface{})
	json.Unmarshal(resp.Body.Bytes(), &res)
	fmt.Println(res)

	assert.Equal(t, 200, resp.Result().StatusCode)
}
