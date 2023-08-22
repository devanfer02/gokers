package configs

import (
	"fmt"
	"os"

	"github.com/devanfer02/gokers/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB 
}

func (baseDb *Database) ConnectToDB() {
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

	baseDb.db = db
}

func (baseDb *Database) MigrateDB() {
	if (baseDb.db == nil) {
		panic(fmt.Errorf("ERROR MIGRATING DB"))
	}

	baseDb.migrate(
		&models.Student{},
		&models.Course{},
		&models.Class{},
		&models.KRS{},
		&models.Lecturer{},
		&models.Admin{},
	)
}

func (baseDb *Database) Create(data interface{}) error {
	return baseDb.db.Model(data).Create(data).Error
}

func (baseDb *Database) FindAll(data interface{}) error {
	return baseDb.db.Model(data).Find(data).Error
}

func (baseDb *Database) Find(query string, data interface{}, params ...interface{}) error {
	return baseDb.db.Model(data).Where(query, params).Find(data).Error
}

func (baseDb *Database) FindFirst(query string, data interface{}, params ...interface{}) error {
	return baseDb.db.Model(data).First(data, query, params).Error
}

func (baseDb *Database) FirstByPK(data interface{}, param interface{}) error {
	return baseDb.db.Model(data).First(data, "id = ?", param).Error
}

func (baseDb *Database) Update(query string, data interface{}, params ...interface{}) int64 {
	return baseDb.db.Model(data).Where(query, params).Updates(data).RowsAffected
}

func (baseDb *Database) Delete(query string, data interface{}, params ...interface{}) int64 {
	return baseDb.db.Unscoped().Where(query, params).Delete(data).RowsAffected
}

func (baseDb *Database) PreloadByPK(foreigns []string, data interface{}, params interface{}) error {
	query := baseDb.db 

	for _, foreign := range foreigns {
		query = query.Preload(foreign)
	}

	return query.Where("id = ?", params).First(data).Error
}

func (baseDb *Database) PreloadByCondition(foreigns []string, data interface{}, condition string, params ...interface{}) error {
	query := baseDb.db 

	for _, foreign := range foreigns {
		query = query.Preload(foreign)
	}

	return query.Where(condition, params...).First(data).Error
}

func (baseDb *Database) Count(query string, model interface{}, params ...string) int64 {
	count := int64(0)

	baseDb.db.
		Model(&model).
		Where(query, params).
		Count(&count)

	return count
}

func (baseDb *Database) migrate(models ...interface{}) {
	for _, model := range models {
		if err := baseDb.db.AutoMigrate(model); err != nil {
			panic(err)
		}
	}
}
