package database

import (
	"fmt"
<<<<<<< HEAD
	"github.com/bimbimprasetyoafif/km/config"
	"github.com/bimbimprasetyoafif/km/model"
=======
	"praktikum/config"
	"praktikum/model"
>>>>>>> 2bbdc36abb92c50f1f5e872d14d94c34757f4f37

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	err error
)

func InitDB(conf config.Config) {

	conectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8",
		conf.DB_USERNAME,
		conf.DB_PASSWORD,
		conf.DB_HOST,
		conf.DB_PORT,
		conf.DB_NAME,
	)

	DB, err = gorm.Open(mysql.Open(conectionString))
	if err != nil {
		fmt.Println("error open conection : ", err)
	}

	DB.AutoMigrate(&model.User{})
}