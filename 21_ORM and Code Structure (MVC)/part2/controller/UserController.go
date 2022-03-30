package controller

import (
	"net/http"
	"part2/config"
	"part2/model"
	"strconv"

	"github.com/labstack/echo/v4"
)

// get all users
func GetUsersController(c echo.Context) error {
	var users []model.User

	if err := config.DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	// your solution here
	user := model.User{}

	id, _ := strconv.Atoi(c.Param("id"))
	if err := config.DB.Find(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if user.ID == 0 {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"messages": "invalid id",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get user",
		"user":     user,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	user := model.User{}
	c.Bind(&user)

	if err := config.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	// your solution here
	user := model.User{}
	id, _ := strconv.Atoi(c.Param("id"))
	if err := config.DB.Delete(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if user.ID == 0 {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"messages": "invalid id",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user",
		"user":    user,
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	// your solution here
	user := model.User{}
	id, _ := strconv.Atoi(c.Param("id"))
	if err := config.DB.First(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if user.ID == 0 {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"messages": "invalid id",
		})
	}
	c.Bind(&user)
	if err := config.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user",
		"user":    user,
	})
}
