package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Priyan-Fadhil-Supriyadi/mini-project/domain"
	"github.com/Priyan-Fadhil-Supriyadi/mini-project/model"
	"github.com/labstack/echo/v4"
)

type CarController struct {
	svc domain.AdapterCarService
}

func (ce *CarController) CreateCarController(c echo.Context) error {
	car := model.Car{}
	c.Bind(&car)

	err := ce.svc.CreateCarService(car)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages": "success",
		"cars":     car,
	})
}

func (ce *CarController) UpdateCarController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	car := model.Car{}
	c.Bind(&car)

	err := ce.svc.UpdateCarService(intID, car)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no change",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "edited",
	})
}

func (ce *CarController) DeleteCarController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	err := ce.svc.DeleteCarByID(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusNoContent, map[string]interface{}{
		"messages": "deleted",
	})
}

func (ce *CarController) GetCarController(c echo.Context) error {
	fmt.Println("eksekusi handler")
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {

	}

	res, err := ce.svc.GetCarByID(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"cars":     res,
	})
}

func (ce *CarController) GetCarsController(c echo.Context) error {
	cars := ce.svc.GetAllCarService()

	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"cars":     cars,
	}, "  ")
}
