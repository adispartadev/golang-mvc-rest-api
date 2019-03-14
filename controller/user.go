package controller

import (
	e "golang-mvc-rest-api/entity"
	m "golang-mvc-rest-api/middleware"
	"golang-mvc-rest-api/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SignUp(c echo.Context) error {

	var user e.User
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	hash, err := m.HashPassword(user.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	user.Password = hash

	err = model.InsertUser(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	token, err := m.Encode(user.Username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, e.SetResponse(http.StatusOK, "success", token))

}

func Login(c echo.Context) error {

	var user e.User
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	userData, err := model.GetUserByUsername(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	isPasswordValid := m.CheckPasswordHash(user.Password, userData.Password)
	if !isPasswordValid {
		return c.JSON(http.StatusBadRequest, "password is not valid")
	}

	token, err := m.Encode(userData.Username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, e.SetResponse(http.StatusOK, "success", token))
}
