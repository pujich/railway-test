package helpers

import (
	"errors"
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
	// "github.com/labstack/echo"
	"github.com/labstack/echo/v4"
)

var secretKey = "apa aja"

func GenerateToken(id uint, email string) string {
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
	}
	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := parseToken.SignedString([]byte(secretKey))

	return signedToken
}

func VerifyToken(c echo.Context) (jwt.MapClaims, error) {
	errResponse := errors.New("sign in to proceed")
	headerToken := c.Request().Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")

	if !bearer {
		return nil, errResponse
	}

	stringToken := strings.Split(headerToken, " ")[1]
	token, _ := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResponse
		}
		return []byte(secretKey), nil
	})

	fmt.Printf("INI TOKEN STRING %v\n", stringToken)
	fmt.Printf("INI TOKEN SEMUA %v", token)

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, errResponse
	}
	return token.Claims.(jwt.MapClaims), nil
}
