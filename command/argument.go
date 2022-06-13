package command

import (
	"os"
)

func Arguments() {
	if len(os.Args) == 2 || len(os.Args) == 3 {
		argument := os.Args[1:]
		if argument[0] == Migrate {
			ArgMigrate()
		} else if argument[0] == Seeder {
			ArgSeeder()
		}
	}
}
