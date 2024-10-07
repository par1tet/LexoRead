package initdb

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	DB *gorm.DB
}

func Init(dsn string) (*DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})
	if err != nil {
		return nil, err
	}
	return &DB{DB: db}, nil
}

func (db *DB) Migrate(models ...interface{}) error {
	err := db.DB.AutoMigrate(models...)
	if err != nil {
		return err
	}
	return nil
}
