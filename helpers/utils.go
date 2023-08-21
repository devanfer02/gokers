package helpers

import (
	"fmt"
	"time"

	"github.com/devanfer02/gokers/configs"
	"github.com/devanfer02/gokers/models"
	"github.com/google/uuid"
)

func CreateNIM(major, faculty, entrance string, Db *configs.Database) (string, error) {
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
	lastCode     := determineLastCode("major = ? AND entrance = ?", models.Student{}, Db, major, entrance)

	nim := year + facultyCode + majorCode + entranceCode + lastCode

	return nim, nil
}

func CreateNDN(major, faculty string, Db *configs.Database) (string, error) {
	yearCode 	:= fmt.Sprintf("%02d", (time.Now().Year() % 100))

	facultyCode := configs.FacultyMap[faculty]["code"]
	
	if _, ok := configs.FacultyMap[faculty][major]; !ok {
		return "", fmt.Errorf("major code doesnt exist, major: %s", major)
	}

	majorCode := configs.FacultyMap[faculty][major]
	lastCode  := determineLastCode("major = ?", models.Lecturer{}, Db, major)

	nim := "00" + yearCode + facultyCode + majorCode + lastCode

	return nim, nil
}

func GenerateUUID() uuid.UUID {
	newUUID, _ := uuid.NewRandom()

	return newUUID
}

func determineLastCode(query string, model interface{},  Db *configs.Database, params ...string) string {
	count := Db.Count(query, model, params...)

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

