package middleware

import (
	"errors"
	"fmt"
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
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256

type JWTClaim struct {
	jwt.StandardClaims
	Username string `json:"username"`
}

func Encode(username string) (string, error) {
	claims := JWTClaim{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			Issuer:    APPLICATION_NAME,
			ExpiresAt: time.Now().Add(LOGIN_EXPIRATION_DURATION).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(JWT_SIGNING_METHOD, claims)

	t, err := token.SignedString([]byte(JWT_SECRET_KEY))
	if err != nil {
		return "", err
	}
	return t, nil
}

func IsAuthorized(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusBadRequest, "authorization header is required")
		}

		if !strings.Contains(authHeader, "Bearer") {
			return c.JSON(http.StatusBadRequest, "authorization token is invalid")
		}

		tokenString1 := strings.Replace(authHeader, "Bearer", " ", -1)
		tokenString2 := strings.TrimSpace(tokenString1)

		fmt.Println("======token string=======")
		fmt.Printf("%+v\n", tokenString2)
		fmt.Println("=============")

		token, err := jwt.Parse(tokenString2, func(token *jwt.Token) (interface{}, error) {
			method, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, errors.New("signing method is invalid")
			}
			if method != JWT_SIGNING_METHOD {
				return nil, errors.New("signing method is invalid")
			}
			return []byte(JWT_SECRET_KEY), nil
		})

		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			return c.JSON(http.StatusBadRequest, "token payload is invalid")
		}
		c.Set("userInfo", claims)
		return next(c)
	}
}
