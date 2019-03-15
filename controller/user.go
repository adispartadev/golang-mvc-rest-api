package controller

import (
	e "golang-mvc-rest-api/entity"
	m "golang-mvc-rest-api/middleware"
	"golang-mvc-rest-api/model"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
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

	token, err := m.Encode(user.Username, m.LOGIN_EXPIRATION_DURATION, os.Getenv("JWT_SECRET_KEY"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	refreshToken, err := m.Encode(user.Username, m.REFRESH_TOKEN_DURATION, os.Getenv("JWT_REFRESH_KEY"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = model.UpdateRefreshToken(&user, refreshToken)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, e.SetTokenResponse(http.StatusOK, "success", token, refreshToken))

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

	token, err := m.Encode(userData.Username, m.LOGIN_EXPIRATION_DURATION, os.Getenv("JWT_SECRET_KEY"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	refreshToken, err := m.Encode(userData.Username, m.REFRESH_TOKEN_DURATION, os.Getenv("JWT_REFRESH_KEY"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = model.UpdateRefreshToken(&user, refreshToken)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, e.SetTokenResponse(http.StatusOK, "success", token, refreshToken))
}

func RefreshToken(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return c.JSON(http.StatusBadRequest, "authorization header is required")
	}

	if !strings.Contains(authHeader, "Bearer") {
		return c.JSON(http.StatusBadRequest, "authorization token is invalid")
	}

	tokenString1 := strings.Replace(authHeader, "Bearer", " ", -1)
	tokenString2 := strings.TrimSpace(tokenString1)

	claims := m.JWTClaim{}

	token, err := jwt.ParseWithClaims(tokenString2, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_REFRESH_KEY")), nil
	})

	err = token.Claims.Valid()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err == jwt.ErrSignatureInvalid {
		return c.JSON(http.StatusUnauthorized, err.Error())
	}

	if err != nil {
		return c.JSON(http.StatusUnauthorized, err.Error())
	}

	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		return c.JSON(http.StatusBadRequest, "token is still valid")
	}

	expirationTime := m.LOGIN_EXPIRATION_DURATION
	claims.ExpiresAt = expirationTime
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := newToken.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, t)

}
