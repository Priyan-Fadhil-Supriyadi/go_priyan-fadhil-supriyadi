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

func TestCreateRentControllerAll(t *testing.T) {
	svc := MockRentSvc{}

	svc.On("CreateRentService", mock.Anything).
		Return(errors.New("new")).Once()

	svc.On("CreateRentService", mock.Anything).
		Return(nil).Once()

	usrController := RentController{
		svc: &svc,
	}

	e := echo.New()

	t.Run("error", func(t *testing.T) {
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		usrController.CreateRentController(echoContext)

		assert.Equal(t, 500, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		usrController.CreateRentController(echoContext)

		assert.Equal(t, 201, w.Result().StatusCode)
	})
}

func TestGetRentsControllerAll(t *testing.T) {
	svc := MockRentSvc{}

	svc.On("GetAllRentService", mock.Anything).
		Return(errors.New("new")).Once()

	svc.On("GetAllRentService", mock.Anything).
		Return(nil).Once()

	usrController := RentController{
		svc: &svc,
	}

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		usrController.GetRentsController(echoContext)

		assert.Equal(t, 200, w.Result().StatusCode)
	})
}

func TestDeleteRentControllerAll(t *testing.T) {
	svc := MockRentSvc{}

	svc.On("DeleteRentService", mock.Anything).
		Return(nil).Once()

	usrController := RentController{
		svc: &svc,
	}

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("DELETE", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		usrController.DeleteRentController(echoContext)

		assert.Equal(t, 204, w.Result().StatusCode)
	})
}
