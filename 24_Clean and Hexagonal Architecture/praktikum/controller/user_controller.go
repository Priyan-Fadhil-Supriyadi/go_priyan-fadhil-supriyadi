package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"praktikum/domain"
	"praktikum/model"
)

type controllerEcho struct {
	repo domain.AdapterRepository
}

func (ce *controllerEcho) CreateUserController(c echo.Context) error {
	user := model.User{}
	c.Bind(&user)

	err := ce.repo.CreateUsers(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages": "success",
		"users":    user,
	})
}

func (ce *controllerEcho) UpdateUserController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	user := model.User{}
	c.Bind(&user)

	err := ce.repo.UpdateOneByID(intID, user)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no change",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "edited",
	})
}

func (ce *controllerEcho) DeleteUserController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	err := ce.repo.DeleteByID(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusNoContent, map[string]interface{}{
		"messages": "deleted",
	})
}

func (ce *controllerEcho) GetUserController(c echo.Context) error {
	fmt.Println("eksekusi handler")
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {

	}

	res, err := ce.repo.GetOneByID(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"users":    res,
	})
}

func (ce *controllerEcho) GetUsersController(c echo.Context) error {
	users := ce.repo.GetAll()

	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"users":    users,
	}, "  ")
}

func NewEchoController(repo domain.AdapterRepository) domain.AdapterController {
	return &controllerEcho{
		repo: repo,
	}
}
