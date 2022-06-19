package command

import (
	"flag"
	"github.com/urfave/cli"
	"go_harians/database/migrations"
	"go_harians/database/seeders"
	"gorm.io/gorm"
	"log"
	"os"
)

func InitCommands(db *gorm.DB) {
	arg := flag.Arg(1)

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
				err := seeders.DBSeed(db)
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name: "make:table",
			Action: func(c *cli.Context) error {
				MakeMigration(arg)
				MakeDTO(arg)
				MakeRepository(arg)
				return nil
			},
		},
	}
	err := cmdApp.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
