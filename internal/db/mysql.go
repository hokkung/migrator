package db

import (
	"github.com/hokkung/migrator/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB(config config.MysqlProperties) (*gorm.DB, error) {
	dsn := config.DNS()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	err = sqlDB.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
