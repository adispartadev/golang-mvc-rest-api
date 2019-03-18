package controller

import (
	e "golang-mvc-rest-api/entity"
	"golang-mvc-rest-api/model"
	"io"
	"net/http"
	"os"
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
	page, err := strconv.Atoi(c.Param("page"))
	limit, err := strconv.Atoi(c.Param("limit"))

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

	return c.JSON(http.StatusCreated, e.SetResponse(http.StatusCreated, "ok", EmptyValue))

}

func RemoveOwner(c echo.Context) error {
	ownerID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, e.SetResponse(http.StatusBadRequest, err.Error(), EmptyValue))
	}

	err = model.DeleteOwner(ownerID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, e.SetResponse(http.StatusBadRequest, err.Error(), EmptyValue))
	}

	return c.JSON(http.StatusAccepted, "ok")
}

func EditOwner(c echo.Context) error {
	var owner e.Owner
	ownerID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, e.SetResponse(http.StatusBadRequest, err.Error(), EmptyValue))
	}

	err = c.Bind(&owner)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, e.SetResponse(http.StatusUnprocessableEntity, err.Error(), EmptyValue))
	}

	owner.ID = ownerID

	err = model.UpdateOwner(&owner)
	if err != nil {
		return c.JSON(http.StatusBadRequest, e.SetResponse(http.StatusBadRequest, err.Error(), EmptyValue))
	}

	return c.JSON(http.StatusOK, e.SetResponse(http.StatusOK, "edited", EmptyValue))
}

func AddOwnerImage(c echo.Context) error {

	maxFileSize := 1000000

	file, err := c.FormFile("owner_profile")
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if file.Header["Content-Type"][0] != "image/png" {
		return c.JSON(http.StatusBadRequest, "only able to upload png file")
	}

	if file.Size > int64(maxFileSize) {
		return c.JSON(http.StatusBadRequest, "file size can't be more than 1 MB")
	}

	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	defer src.Close()

	uploadFilePath := "./tmp/" + file.Filename

	dst, err := os.Create(uploadFilePath)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, "ok")
}
