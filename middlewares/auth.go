package middlewares

import (
	"net/http"
	"time"

	"github.com/devanfer02/gokers/configs"
	"github.com/devanfer02/gokers/helpers"
	"github.com/devanfer02/gokers/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AuthMiddleware struct {
	Db *configs.Database
}

func (authMdl *AuthMiddleware) RequireAuth(ctx *gin.Context) {
	tokenStr, err := ctx.Cookie("Authorization")

	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return 
	}

	claims, ok, token := helpers.CreateMapClaims(tokenStr)

	if !ok || !token.Valid {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return 
	}

	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return 
	}

	var student models.Student

	authMdl.Db.FirstByPK(&student, claims["sub"])

	if stdID, err := uuid.Parse(student.ID.String()); err != nil || stdID == uuid.Nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return 
	}

	ctx.Set("std", student)
	ctx.Next()
}

func (authMdl *AuthMiddleware) RequireAdmin(ctx *gin.Context) {
	tokenStr, err := ctx.Cookie("AdminAuthorization")

	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

	claims, ok, token := helpers.CreateMapClaims(tokenStr)

	if !ok || !token.Valid {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return 
	}

	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return 
	}

	var admin models.Admin

	authMdl.Db.FirstByPK(&admin, claims["adm"])

	if admID, err := uuid.Parse(admin.ID.String()); err != nil || admID == uuid.Nil {
		// berhenti disini
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return 
	}

	ctx.Set("adm", admin)
	ctx.Next()
}