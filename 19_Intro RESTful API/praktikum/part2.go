package main

import (

"net/http" 
"strconv"

"github.com/labstack/echo"

)

type User struct {
	Id int `json:"id" form:"id"`
	Name string `json: "name" form: "name"`
	Email string `json: "email" form:"email" Password string 'json:"password" form:"password"`
}

var users []User

controller

// get all users

func GetUsersController (c echo.Context) error { return c.JSON(http.StatusOK, map[string]interface(){

"messages": "success get all users",

})

"users": users,

get user by id

func GetUserController(c echo.Context) error { // your solution here

}

delete user by id func DeleteUserController (c echo.Context) error {

// your solution here

}

update user by id func UpdateUser Controller (c echo.Context) error ( 
// your solution here