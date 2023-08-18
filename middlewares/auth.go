package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/devanfer02/gokers/configs"
	"github.com/devanfer02/gokers/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/golang-jwt/jwt/v4"
)

func RequireAuth(ctx *gin.Context) {
	tokenStr, err := ctx.Cookie("Authorization")

	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return 
	}

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}

		return []byte(os.Getenv("SECRET_KEY")), nil 
	})

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return 
	}

	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return 
	}

	var student models.Student

	configs.DB.First(&student, claims["sub"])

	if stdID, err := uuid.Parse(student.ID.String()); err != nil || stdID == uuid.Nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return 
	}

	ctx.Set("std", student)

	ctx.Next()
}