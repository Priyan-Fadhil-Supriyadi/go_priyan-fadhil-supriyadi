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

func TestCreateCarControllerAll(t *testing.T) {
	svc := MockCarSvc{}

	svc.On("CreateCarService", mock.Anything).
		Return(errors.New("new")).Once()

	svc.On("CreateCarService", mock.Anything).
		Return(nil).Once()

	usrController := CarController{
		svc: &svc,
	}

	e := echo.New()

	t.Run("error", func(t *testing.T) {
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		usrController.CreateCarController(echoContext)

		assert.Equal(t, 500, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		usrController.CreateCarController(echoContext)

		assert.Equal(t, 201, w.Result().StatusCode)
	})
}

func TestGetCarsControllerAll(t *testing.T) {
	svc := MockCarSvc{}

	svc.On("GetAllCarService", mock.Anything).
		Return(errors.New("new")).Once()

	svc.On("GetAllCarService", mock.Anything).
		Return(nil).Once()

	usrController := CarController{
		svc: &svc,
	}

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		usrController.GetCarsController(echoContext)

		assert.Equal(t, 200, w.Result().StatusCode)
	})
}

func TestDeleteCarControllerAll(t *testing.T) {
	svc := MockCarSvc{}

	svc.On("DeleteCarService", mock.Anything).
		Return(nil).Once()

	usrController := CarController{
		svc: &svc,
	}

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("DELETE", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		usrController.DeleteCarController(echoContext)

		assert.Equal(t, 204, w.Result().StatusCode)
	})
}
