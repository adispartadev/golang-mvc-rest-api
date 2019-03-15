package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

var APPLICATION_NAME = "golang-mvc-rest-api"
var LOGIN_EXPIRATION_DURATION int64
var REFRESH_TOKEN_DURATION int64

type JWTClaim struct {
	jwt.StandardClaims
	Username string `json:"username"`
}

func init() {
	LOGIN_EXPIRATION_DURATION = time.Now().Add(100 * time.Second).Unix()
	REFRESH_TOKEN_DURATION = time.Now().Add(10 * time.Minute).Unix()
}

func Encode(username string, duration int64, secretKey string) (string, error) {

	fmt.Println("======secret key=======")
	fmt.Printf("%+v\n", secretKey)
	fmt.Printf("%+v\n", duration)
	fmt.Println("=============")

	claims := JWTClaim{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			Issuer:    APPLICATION_NAME,
			ExpiresAt: duration,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(secretKey))
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

		token, err := jwt.Parse(tokenString2, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})

		if !token.Valid {
			return c.JSON(http.StatusUnauthorized, "token is not valid")
		}

		if err == jwt.ErrSignatureInvalid {
			return c.JSON(http.StatusUnauthorized, err.Error())
		}

		if err != nil {
			return c.JSON(http.StatusUnauthorized, err.Error())
		}

		method, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return c.JSON(http.StatusBadRequest, "signing method is invalid")
		}
		if method != jwt.SigningMethodHS256 {
			return c.JSON(http.StatusBadRequest, "signing method is invalid")
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.JSON(http.StatusBadRequest, "token payload is invalid")
		}
		c.Set("userInfo", claims)
		return next(c)
	}
}
