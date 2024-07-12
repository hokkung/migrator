package cmd

import (
	"github.com/hokkung/migrator/config"
	"github.com/hokkung/migrator/internal/db"
	"github.com/hokkung/migrator/internal/migrator"
	"github.com/spf13/cobra"
)

var migrateUpCmd = &cobra.Command{
	Use:     "up",
	Short:   "Migrate the database to the latest version",
	Example: `migrator up`,
	Run:     migrateDBUp,
}

func migrateDBUp(cmd *cobra.Command, args []string) {
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

	err = m.Up()
	if err != nil {
		panic(err)
	}
}
