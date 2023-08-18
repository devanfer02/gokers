package helpers

import (
	"fmt"
	"time"

	"github.com/devanfer02/gokers/configs"
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

func CreateNIM(major, faculty, entrance string) (string, error) {
	currentYear := time.Now().Year()
	twodigits	:= currentYear % 100
	year		:= fmt.Sprintf("%02d", twodigits)

	if _, ok := configs.FacultyMap[faculty]["code"]; !ok {
		return "", fmt.Errorf("faculty code doesnt exist, faculty: %s", faculty)
	}

	facultyCode := configs.FacultyMap[faculty]["code"]

	if _, ok := configs.FacultyMap[faculty][major]; !ok {
		return "", fmt.Errorf("major code doesnt exist, major: %s", major)
	}

	majorCode 	 := configs.FacultyMap[faculty][major]
	entranceCode := determineEntranceCode(entrance)
	lastCode     := determineLastCode(major, entrance)

	nim := year + facultyCode + majorCode + entranceCode + lastCode

	return nim, nil
}

func determineLastCode(major, entrance string) string {
	var student models.Student
	count := int64(0)

	configs.DB.
		Model(&student).
		Where("major = ? AND entrance = ?", major, entrance).
		Count(&count)

	if count < 10 {
		return fmt.Sprintf("00%d", count)
	}
	if count < 100 {
		return fmt.Sprintf("0%d", count)
	}
	return fmt.Sprintf("%d", count)
}

func determineEntranceCode(entrance string) string {
	if (entrance == "snbp") {
		return "01111" 
	}  
	if (entrance == "snbt") {
		return "00111"
	} 
		
	return "07111"
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