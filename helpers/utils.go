package helpers

import (
	"fmt"
	"time"
	"unicode"

	"github.com/devanfer02/gokers/configs"
	"github.com/devanfer02/gokers/models"
	"github.com/google/uuid"
)

func GetTotalSks(krs []models.KRS) uint {
	sum := uint(0)
	for _, class := range krs {
		sum += class.Sks 
	}

	return sum 
}

func CheckTypeExist(types, faculty, major *string) error {
	if *types != "university" && *types != "faculty" && *types != "major" {
		return fmt.Errorf("types must at least be university or faculty or major, requested types: %s", *types)
	}

	if *types == "university" {
		*faculty = ""
		*major = ""
		return nil
	}

	if *types == "faculty" {
		if !isAllUpper(*faculty) {
			return fmt.Errorf("faculty must be all uppercase")
		}
		if _, ok := configs.FacultyMap[*faculty]; !ok {
			return fmt.Errorf("faculty doesnt exist, requested faculty: %s", *faculty)
		}
		return nil
	}

	if *types == "major" {
		if !isAllUpper(*faculty) {
			return fmt.Errorf("faculty must be all uppercase")
		}

		facultyMap, ok := configs.FacultyMap[*faculty]

		if !ok {
			return fmt.Errorf("faculty doesnt exist, requested faculty: %s", *faculty)
		}

		if _, ok := facultyMap[*major]; !ok {
			return fmt.Errorf("major doesnt exist, requested major: %s", *major)
		}

		return nil
	}

	return nil
}

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

func isAllUpper(s string) bool {
	for _, char := range s {
		if !unicode.IsUpper(char) {
			return false
		}
	}

	return true
}