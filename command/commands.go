package command

import (
	"github.com/urfave/cli"
	"go_harians/database/migrations"
	"go_harians/database/seeders"
	"gorm.io/gorm"
	"log"
	"os"
)

func InitCommands(db *gorm.DB) {
	cmdApp := cli.NewApp()
	cmdApp.Commands = []cli.Command{
		{
			Name: "db:migrate",
			Action: func(c *cli.Context) error {
				migrations.DBMigrate(db)
				return nil
			},
		},
		{
			Name: "db:seeder",
			Action: func(c *cli.Context) error {
				seeders.DBSeed(db)
				return nil
			},
		},
	}
	err := cmdApp.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
