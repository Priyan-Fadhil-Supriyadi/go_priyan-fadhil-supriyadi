package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Priyan-Fadhil-Supriyadi/mini-project/domain"
	"github.com/Priyan-Fadhil-Supriyadi/mini-project/model"
	"github.com/labstack/echo/v4"
)

type DepotController struct {
	svc domain.AdapterDepotService
}

func (ce *DepotController) CreateDepotController(c echo.Context) error {
	depot := model.Depot{}
	c.Bind(&depot)

	err := ce.svc.CreateDepotService(depot)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages": "success",
		"depots":   depot,
	})
}

func (ce *DepotController) UpdateDepotController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	depot := model.Depot{}
	c.Bind(&depot)

	err := ce.svc.UpdateDepotService(intID, depot)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no change",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "edited",
	})
}

func (ce *DepotController) DeleteDepotController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	err := ce.svc.DeleteDepotByID(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusNoContent, map[string]interface{}{
		"messages": "deleted",
	})
}

func (ce *DepotController) GetDepotController(c echo.Context) error {
	fmt.Println("eksekusi handler")
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {

	}

	res, err := ce.svc.GetDepotByID(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"depots":   res,
	})
}

func (ce *DepotController) GetDepotsController(c echo.Context) error {
	depots := ce.svc.GetAllDepotService()

	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"depots":   depots,
	}, "  ")
}
