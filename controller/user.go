package controller

import (
	e "golang-mvc-rest-api/entity"
	"golang-mvc-rest-api/model"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

const APPLICATION_NAME = "golang-mvc-rest-api"
const LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Hour

var JWT_SECRET_KEY = os.Getenv("JWT_SECRET_KEY")

type JWTClaim struct {
	jwt.StandardClaims
	Username string `json:"username"`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", nil
	}
	return string(bytes), nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func SignUp(c echo.Context) error {

	var user e.User
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	hash, err := HashPassword(user.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	user.Password = hash

	err = model.InsertUser(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, "ok")

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

	isPasswordValid := CheckPasswordHash(user.Password, userData.Password)
	if !isPasswordValid {
		return c.JSON(http.StatusBadRequest, "password is not valid")
	}

	claims := &JWTClaim{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			Issuer:    APPLICATION_NAME,
			ExpiresAt: time.Now().Add(LOGIN_EXPIRATION_DURATION).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(JWT_SECRET_KEY))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, e.SetResponse(http.StatusOK, "success", t))
}
