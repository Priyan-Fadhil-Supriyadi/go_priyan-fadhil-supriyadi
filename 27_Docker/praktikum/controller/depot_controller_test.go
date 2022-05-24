package controller

import (
	// "errors"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	// "github.com/stretchr/testify/mock"
)

func TestCreateDepotControllerAll(t *testing.T) {
	svc := MockDepotSvc{}

	svc.On("CreateDepotService", mock.Anything).
		Return(errors.New("new")).Once()

	svc.On("CreateDepotService", mock.Anything).
		Return(nil).Once()

	usrController := DepotController{
		svc: &svc,
	}

	e := echo.New()

	t.Run("error", func(t *testing.T) {
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		usrController.CreateDepotController(echoContext)

		assert.Equal(t, 500, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		usrController.CreateDepotController(echoContext)

		assert.Equal(t, 201, w.Result().StatusCode)
	})
}

func TestGetDepotsControllerAll(t *testing.T) {
	svc := MockDepotSvc{}

	svc.On("GetAllDepotService", mock.Anything).
		Return(errors.New("new")).Once()

	svc.On("GetAllDepotService", mock.Anything).
		Return(nil).Once()

	usrController := DepotController{
		svc: &svc,
	}

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		usrController.GetDepotsController(echoContext)

		assert.Equal(t, 200, w.Result().StatusCode)
	})
}

func TestDeleteDepotControllerAll(t *testing.T) {
	svc := MockDepotSvc{}

	svc.On("DeleteDepotService", mock.Anything).
		Return(nil).Once()

	usrController := DepotController{
		svc: &svc,
	}

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("DELETE", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		usrController.DeleteDepotController(echoContext)

		assert.Equal(t, 204, w.Result().StatusCode)
	})
}
