package migrator

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"gorm.io/gorm"
)

func NewMySQLMigrator(gorm *gorm.DB) (*migrate.Migrate, error) {
	engine, err := gorm.DB()
	if err != nil {
		return nil, err
	}

	driver, err := mysql.WithInstance(engine, &mysql.Config{})
	if err != nil {
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance("internal/migrations", "sql", driver)
	if err != nil {
		return nil, err
	}

	return m, nil
}
