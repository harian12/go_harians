package command

import (
	"os"
)

func Arguments() {
	if len(os.Args) > 1 && len(os.Args) < 5 {
		argument := os.Args[1:]

		switch argument[0] {
		case Migrate:
			ArgMigrate()
		case Seeder:
			ArgSeeder()
		case MakeRepository:
			ArgRepository()
		case MakeService:
			ArgService()
		case MakeHandler:
			ArgHandler()
		}
	}
}
