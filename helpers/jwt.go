package helpers

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)


func GetTokenStr(id uuid.UUID, subject string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		subject: id, 
		"exp": time.Now().Add(time.Hour * 12).Unix(),
	})

	tokenStr, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	
	return tokenStr, err 
}

func CreateMapClaims(tokenStr string) (jwt.MapClaims, bool, *jwt.Token ){
	token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}

		return []byte(os.Getenv("SECRET_KEY")), nil 
	})

	claims, ok := token.Claims.(jwt.MapClaims)

	return claims, ok, token
}