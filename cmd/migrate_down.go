package cmd

import (
	"github.com/hokkung/migrator/config"
	"github.com/hokkung/migrator/internal/db"
	"github.com/hokkung/migrator/internal/migrator"
	"github.com/spf13/cobra"
)

var migrateDownCmd = &cobra.Command{
	Use:     "down",
	Short:   "Migrate the database down to the first version",
	Example: `migrator down`,
	Run:     migrateDBDown,
}

func migrateDBDown(cmd *cobra.Command, args []string) {
	appConfig := config.NewAppProperties()
	sqlConfig := config.NewMysqlProperties()
	_ = config.NewConfiguration(*appConfig, *sqlConfig)
	gorm, err := db.NewDB(*sqlConfig)
	if err != nil {
		panic(err)
	}

	m, err := migrator.NewMySQLMigrator(gorm)
	if err != nil {
		panic(err)
	}

	err = m.Down()
	if err != nil {
		panic(err)
	}
}
