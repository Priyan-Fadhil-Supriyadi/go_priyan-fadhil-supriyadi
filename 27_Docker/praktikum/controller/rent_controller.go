package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Priyan-Fadhil-Supriyadi/mini-project/domain"
	"github.com/Priyan-Fadhil-Supriyadi/mini-project/model"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type RentController struct {
	svc domain.AdapterRentService
}

func (ce *RentController) CreateRentController(c echo.Context) error {
	rent := model.Rent{}
	c.Bind(&rent)

	err := ce.svc.CreateRentService(rent)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages": "success",
		"rents":    rent,
	})
}

func (ce *RentController) UpdateRentController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	userID := c.Param("userid")
	userIntID, _ := strconv.Atoi(userID)

	user := model.User{}
	c.Bind(&user)

	bearer := c.Get("user").(*jwt.Token)
	claim := bearer.Claims.(jwt.MapClaims)

	rent := model.Rent{}
	c.Bind(&rent)

	err := ce.svc.UpdateRentService(intID, userIntID, int(claim["id"].(float64)), rent)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no change or unauthorization",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages":   "edited",
		"id":         intID,
		"expeted id": claim["id"],
	})
}

func (ce *RentController) DeleteRentController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	err := ce.svc.DeleteRentByID(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusNoContent, map[string]interface{}{
		"messages": "deleted",
	})
}

func (ce *RentController) GetRentController(c echo.Context) error {
	fmt.Println("eksekusi handler")
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {

	}

	res, err := ce.svc.GetRentByID(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"rents":    res,
	})
}

func (ce *RentController) GetRentsController(c echo.Context) error {
	rents := ce.svc.GetAllRentService()

	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"rents":    rents,
	}, "  ")
}
