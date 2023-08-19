package helpers

import (
	"fmt"
	"os"
	"time"

	"github.com/devanfer02/gokers/configs"
	"github.com/devanfer02/gokers/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
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
	currentYear := time.Now().Year() % 100
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
	lastCode     := determineLastCode("major = ? AND entrance = ?", models.Student{}, major, entrance)

	nim := year + facultyCode + majorCode + entranceCode + lastCode

	return nim, nil
}

func CreateNDN(major, faculty string) (string, error) {
	yearCode 	:= fmt.Sprintf("%02d", (time.Now().Year() % 100))

	facultyCode := configs.FacultyMap[faculty]["code"]
	
	if _, ok := configs.FacultyMap[faculty][major]; !ok {
		return "", fmt.Errorf("major code doesnt exist, major: %s", major)
	}

	majorCode := configs.FacultyMap[faculty][major]
	lastCode  := determineLastCode("major = ?", models.Lecturer{}, major)

	nim := "00" + yearCode + facultyCode + majorCode + lastCode

	return nim, nil

}

func GetTokenStr(id uuid.UUID) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": id, 
		"exp": time.Now().Add(time.Hour * 12).Unix(),
	})

	tokenStr, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		return "", err
	}
	return tokenStr, nil
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

func determineLastCode(query string, model interface{}, params ...string) string {
	count := DbCount(query, model, params...)

	if count++; count < 10 {
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

func DbCount(query string, model interface{}, params ...string) int64 {
	count := int64(0)

	configs.DB.
		Model(&model).
		Where(query, params).
		Count(&count)

	return count
}
