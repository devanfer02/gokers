package helpers

import (
	"github.com/devanfer02/gokers/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(hashed), err
}

func CompPassword(db, body *string) error {
	return bcrypt.CompareHashAndPassword([]byte(*db), []byte(*body))
}

func GetStudentID(ctx *gin.Context) (uuid.UUID, bool) {
	student, _ := ctx.Get("std")

	switch u := student.(type) {
		case models.Student :
			return u.ID, true 
		default :
			return uuid.Nil, false
	}
}

func GenerateUUID() uuid.UUID {
	newUUID, _ := uuid.NewRandom()

	return newUUID
}