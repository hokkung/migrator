package cmd

import (
	"errors"
	"github.com/hokkung/migrator/config"
	"github.com/hokkung/migrator/internal/db"
	"github.com/hokkung/migrator/internal/migrator"
	"github.com/spf13/cobra"
	"strconv"
)

var migrateStepCmd = &cobra.Command{
	Use:     "step",
	Short:   "Migrate the database up/down to the given input",
	Example: `migrator step <step>`,
	Run:     migrateDBStep,
	Args:    valiadateStepArg,
}

func valiadateStepArg(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.New("you must specify a step")
	}
	step, err := strconv.Atoi(args[0])
	if err != nil {
		panic(err)
	}
	if step == 0 {
		return errors.New("step must be greater than 0 or less than 0")
	}
	return nil
}

func migrateDBStep(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		panic("Must specify a step to migrate")
	}

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

	step, err := strconv.Atoi(args[0])
	if err != nil {
		panic(err)
	}

	err = m.Steps(step)
	if err != nil {
		panic(err)
	}
}
