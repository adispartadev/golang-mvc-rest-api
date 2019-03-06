package controller

import (
	"golang-mvc-rest-api/model"
	"net/http"
	"regexp"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetOwners(c echo.Context) error {

	res, err := model.GetAllOwners()
	if err != nil {
		return c.JSON(http.StatusBadRequest, SetResponse(http.StatusBadRequest, err.Error(), EmptyValue))
	}

	if len(res) == 0 {
		return c.JSON(http.StatusOK, SetResponse(http.StatusOK, "owners is empty", EmptyValue))
	}

	return c.JSON(http.StatusOK, SetResponse(http.StatusOK, "", res))
}

func GetOwnersLimit(c echo.Context) error {
	r := regexp.MustCompile("^[0-9]+$")
	tmpPage := c.Param("page")
	tmpLimit := c.Param("limit")
	if !r.MatchString(tmpPage) {
		return c.JSON(http.StatusBadRequest, "check your payload for page")
	}
	if !r.MatchString(tmpLimit) {
		return c.JSON(http.StatusBadRequest, "check your payload for limit")
	}

	page, err := strconv.Atoi(tmpPage)
	limit, err := strconv.Atoi(tmpLimit)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "please check page and limit value")
	}

	totalOwner, err := model.CountOwners()
	if err != nil {
		return c.JSON(http.StatusBadRequest, SetResponse(http.StatusBadRequest, err.Error(), EmptyValue))
	}

	pagination := HandlePagination(page, limit, totalOwner)
	owners, err := model.GetAllOwners(pagination)
	if err != nil {
		return c.JSON(http.StatusBadRequest, SetResponse(http.StatusBadRequest, err.Error(), EmptyValue))
	}

	return c.JSON(http.StatusOK, owners)
}
