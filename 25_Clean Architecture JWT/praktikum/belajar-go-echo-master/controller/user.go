package controller

import (
	"belajar-go-echo/model"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetAllUsers(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		users := make([]model.User, 0)
		err := db.Find(&users).Error
		if err != nil {
			return c.JSON(500, echo.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(200, echo.Map{
			"data": users,
		})
	}
}

func CreateUser(db *gorm.DB) echo.HandlerFunc {
	user := model.User{}
	return func(c echo.Context) error {
		if err := c.Bind(&user); err != nil {
			return c.JSON(400, echo.Map{
				"error": err.Error(),
			})
		}
		err := db.Create(&user).Error
		if err != nil {
			return c.JSON(500, echo.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(200, echo.Map{
			"data": user,
		})
	}
}

func Login(db *gorm.DB) echo.HandlerFunc {
	user := model.User{}

	return func(c echo.Context) error {
		err := c.Bind(&user)

		if err != nil {
			return c.JSON(http.StatusUnprocessableEntity, err.Error())
		}

		if user.Username != "admin" && user.Password != "admin" {
			return c.JSON(http.StatusUnauthorized, "username dan password salah")
		}

		claims := jwt.MapClaims{}
		claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		jwt, err := token.SignedString([]byte("secret"))

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, jwt)
	}

}
