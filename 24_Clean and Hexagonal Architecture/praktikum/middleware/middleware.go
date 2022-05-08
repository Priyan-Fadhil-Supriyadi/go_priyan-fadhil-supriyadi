package middleware

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func AuthUser(username, password string, c echo.Context) (bool, error) {
	if username == "admin" && password == "admin" {
		return true, nil
	}

	return false, nil
}
