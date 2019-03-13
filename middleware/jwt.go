package middleware

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

const APPLICATION_NAME = "golang-mvc-rest-api"
const LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Hour

var JWT_SECRET_KEY = os.Getenv("JWT_SECRET_KEY")

type JWTClaim struct {
	jwt.StandardClaims
	Username string `json:"username"`
}

func Decode(username string) (string, error) {
	claims := JWTClaim{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			Issuer:    APPLICATION_NAME,
			ExpiresAt: time.Now().Add(LOGIN_EXPIRATION_DURATION).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(JWT_SECRET_KEY))
	if err != nil {
		return "", err
	}
	return t, nil
}

func isAuthorized(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusBadRequest, "authorization header is required")
		}

		if !strings.Contains(authHeader, "Bearer") {
			return c.JSON(http.StatusBadRequest, "authorization token is invalid")
		}

		tokenString := strings.Replace(authHeader, "Bearer", "", -1)

	}
}
