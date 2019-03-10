package controller

import (
	e "golang-mvc-rest-api/entity"
	"golang-mvc-rest-api/model"
	"net/http"
	"regexp"
	"strconv"

	"github.com/labstack/echo/v4"
)

var (
	EmptyValue = make([]int, 0)
)

// func GetOwners(c echo.Context) error {

// 	res, err := model.GetAllOwners()
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, e.SetResponse(http.StatusBadRequest, err.Error(), EmptyValue))
// 	}

// 	if len(res) == 0 {
// 		return c.JSON(http.StatusOK, e.SetResponse(http.StatusOK, "owners is empty", EmptyValue))
// 	}

// 	return c.JSON(http.StatusOK, e.SetResponse(http.StatusOK, "", res))
// }

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
		return c.JSON(http.StatusBadRequest, e.SetResponse(http.StatusBadRequest, err.Error(), EmptyValue))
	}

	if totalOwner == 0 {
		return c.JSON(http.StatusBadRequest, e.SetResponse(http.StatusBadRequest, err.Error(), EmptyValue))
	}

	pagination := e.HandlePagination(page, limit, totalOwner)
	owners, err := model.GetAllOwners(pagination)
	if err != nil {
		return c.JSON(http.StatusBadRequest, e.SetResponse(http.StatusBadRequest, err.Error(), EmptyValue))
	}

	return c.JSON(http.StatusOK, e.SetPaginationResponse(http.StatusOK, "success", owners, &pagination))
}

func AddOwner(c echo.Context) error {
	var owner e.Owner
	err := c.Bind(&owner)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, e.SetResponse(http.StatusBadRequest, err.Error(), EmptyValue))
	}

	err = model.InsertOwner(&owner)
	if err != nil {
		return c.JSON(http.StatusBadRequest, e.SetResponse(http.StatusBadRequest, err.Error(), EmptyValue))
	}

	return c.JSON(http.StatusCreated, "ok")

}
