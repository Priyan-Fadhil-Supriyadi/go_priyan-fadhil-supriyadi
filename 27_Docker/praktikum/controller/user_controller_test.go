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

func TestCreateUserControllerAll(t *testing.T) {
	svc := MockUserSvc{}

	svc.On("CreateUserService", mock.Anything).
		Return(errors.New("new")).Once()

	svc.On("CreateUserService", mock.Anything).
		Return(nil).Once()

	usrController := EchoController{
		svc: &svc,
	}

	e := echo.New()

	t.Run("error", func(t *testing.T) {
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		usrController.CreateUserController(echoContext)

		assert.Equal(t, 500, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		usrController.CreateUserController(echoContext)

		assert.Equal(t, 201, w.Result().StatusCode)
	})
}

func TestGetUsersControllerAll(t *testing.T) {
	svc := MockUserSvc{}

	svc.On("GetAllUserService", mock.Anything).
		Return(errors.New("new")).Once()

	svc.On("GetAllUserService", mock.Anything).
		Return(nil).Once()

	usrController := EchoController{
		svc: &svc,
	}

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		usrController.GetUsersController(echoContext)

		assert.Equal(t, 200, w.Result().StatusCode)
	})
}

func TestDeleteUserControllerAll(t *testing.T) {
	svc := MockUserSvc{}

	svc.On("DeleteUserService", mock.Anything).
		Return(nil).Once()

	usrController := EchoController{
		svc: &svc,
	}

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("DELETE", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		usrController.DeleteUserController(echoContext)

		assert.Equal(t, 204, w.Result().StatusCode)
	})
}
