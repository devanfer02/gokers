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

func NewDatabase() *Database {
	return &Database{}
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
		&models.Lecturer{},
		&models.Course{},
		&models.Class{},
		&models.KRS{},
		&models.Admin{},
		&models.KrsDetail{},
		&models.CoursePrequisites{},
	)

	baseDb.db.Model(&models.Class{})
}

func (baseDb *Database) Create(data interface{}) error {
	return baseDb.db.Model(data).Create(data).Error
}

func (baseDb *Database) FindAll(data interface{}) error {
	return baseDb.db.Model(data).Find(data).Error
}

func (baseDb *Database) FindAllCondition(query string, data interface{}, params ...interface{}) error {
	return baseDb.db.Model(data).Where(query, params...).Find(data).Error
}

func (baseDb *Database) Find(query string, data interface{}, params ...interface{}) error {
	return baseDb.db.Model(data).Where(query, params...).Find(data).Error
}

func (baseDb *Database) FindFirst(query string, data interface{}, param interface{}) error {
	return baseDb.db.Model(data).First(data, query, param).Error
}

func (baseDb *Database) FirstByPK(data interface{}, param interface{}) error {
	return baseDb.db.Model(data).First(data, "id = ?", param).Error
}

func (baseDb *Database) Update(query string, data interface{}, params ...interface{}) int64 {
	return baseDb.db.Model(data).Where(query, params).Updates(data).RowsAffected
}

func (baseDb *Database) UpdateByPK(data interface{}, param interface{}) int64 {
	return baseDb.db.Model(data).Where("id = ?", param).Updates(data).RowsAffected
}

func (baseDb *Database) UpdateMapByCondition(query string, model interface{}, data map[string]interface{}, params ...interface{}) int64 {
	return baseDb.db.Model(model).Where(query, params).Updates(&data).RowsAffected
}

func (baseDb *Database) Delete(query string, data interface{}, params ...interface{}) int64 {
	return baseDb.db.Unscoped().Where(query, params...).Delete(data).RowsAffected
}

func (baseDb *Database) DeleteByPK(data interface{}, param interface{}) int64 {
	return baseDb.db.Unscoped().Where("id = ?", param).Delete(data).RowsAffected
}


func (baseDb *Database) CountJoins(joins string, queries []string, data interface{}, params []interface{}) int64 {
	count := int64(0)

	db := baseDb.db.Model(data).Joins(joins)

	for i, query := range queries {
		db = db.Where(query, params[i])
	}
	err := db.Count(&count).Error 

	if err != nil {
		panic(err)
	}

	return count
}

func (baseDb *Database) PreloadChainByPK(foreigns []string, data interface{}, params interface{}) error {
	chainedForeign := ""
	for i, foreign := range foreigns {
		chainedForeign += foreign
		if i != len(foreigns) - 1 {
			chainedForeign += "."
		}
	}

	return baseDb.db.Preload(chainedForeign).Where("id = ?", params).First(data).Error; 
}

func (baseDb *Database) PreloadByPK(foreigns []string, data interface{}, params interface{}) error {
	query := baseDb.db 

	for _, foreign := range foreigns {
		query = query.Preload(foreign)
	}

	return query.Where("id = ?", params).First(data).Error
}

func (baseDb *Database) PreloadMany(foreigns []string, data interface{}) error {
	query := baseDb.db 

	for _, foreign := range foreigns {
		query = query.Preload(foreign)
	}

	return query.Find(data).Error
}

func (baseDb *Database) PreloadByCondition(foreigns []string, data interface{}, condition string, params ...interface{}) error {
	query := baseDb.db 

	for _, foreign := range foreigns {
		query = query.Preload(foreign)
	}

	return query.Where(condition, params...).Find(data).Error
}

func (baseDb *Database) Count(query string, model interface{}, params ...interface{}) int64 {
	count := int64(0)

	baseDb.db.
		Model(&model).
		Where(query, params...).
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
