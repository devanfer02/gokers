package helpers

import (
	"strings"

	"github.com/devanfer02/gokers/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetStudentID(ctx *gin.Context) (uuid.UUID, bool) {
	student, _ := ctx.Get("std")

	switch u := student.(type) {
		case models.Student :
			return u.ID, true 
		default :
			return uuid.Nil, false
	}
}

func GetParamID(ctx *gin.Context) (uuid.UUID, error) {
	idParam := ctx.Param("id")
	idParam = strings.ReplaceAll(idParam, "-", "")

	id, err := uuid.Parse(idParam)

	return id, err
}