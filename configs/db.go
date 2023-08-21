package configs

import (
	"fmt"
	"os"

	"github.com/devanfer02/gokers/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Database struct {
	DB *gorm.DB 
}

func (selfs *Database) ConnectToDB() {
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

	selfs.DB = db
	DB = db
}

func (selfs *Database) MigrateDB() {
	if (selfs.DB == nil) {
		fmt.Printf("ERR MIGRATING DB")
		return 
	}

	selfs.migrate(
		&models.Student{},
		&models.Course{},
		&models.Class{},
		&models.KRS{},
		&models.Lecturer{},
	)
}

func (selfs *Database) Create(data interface{}) error {
	return selfs.DB.Create(data).Error
}

func (selfs *Database) FindAll(data interface{}) error {
	return selfs.DB.Find(data).Error
}

func (selfs *Database) Find(query string, data interface{}, params ...interface{}) error {
	return selfs.DB.Model(data).Where(query, params).Find(data).Error
}

func (selfs *Database) FindFirst(query string, data interface{}, params ...interface{}) error {
	return selfs.DB.Model(data).First(data, query, params).Error
}

func (selfs *Database) Update(query string, data interface{}, params ...interface{}) int64 {
	return selfs.DB.Model(data).Where(query, params).Updates(data).RowsAffected
}

func (selfs *Database) Delete(query string, data interface{}, params ...interface{}) int64 {
	return selfs.DB.Unscoped().Where(query, params).Delete(data).RowsAffected
}

func (selfs *Database) Count(query string, model interface{}, params ...string) int64 {
	count := int64(0)

	selfs.DB.
		Model(&model).
		Where(query, params).
		Count(&count)

	return count
}

func (selfs *Database) migrate(models ...interface{}) {
	for _, model := range models {
		if err := selfs.DB.AutoMigrate(model); err != nil {
			panic(err)
		}
	}
}
