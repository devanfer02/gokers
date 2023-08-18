package configs

import (
	"fmt"
	"os"

	"github.com/devanfer02/gokers/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	dsn := fmt.Sprintf(
		"%s%s:@tcp(%s)/%s?parseTime=true", 
		os.Getenv("DB_USERNAME"), 
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_CONNECTION"), 
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(mysql.Open(dsn))

	if err != nil {
		panic(err)
	}

	DB = db
}

func MigrateDB() {
	if (DB == nil) {
		fmt.Printf("ERR MIGRATING DB")
		return 
	}

	migrate(
		&models.Student{},
		&models.Course{},
		&models.Class{},
		&models.KRS{},
		&models.Lecturer{},
	)
}

func migrate(models ...interface{}) {
	for _, model := range models {
		if err := DB.AutoMigrate(model); err != nil {
			panic(err)
		}
	}
}
