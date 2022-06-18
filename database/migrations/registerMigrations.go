package migrations

import (
	"gorm.io/gorm"
	"log"
)

type Migration struct {
	Tabel interface{}
}

func RegisterTable() []Migration {
	return []Migration{
		//daftarkan model atau tabel migration disini
		{Tabel: Users{}},
	}
}

func DBMigrate(db *gorm.DB) {
	for _, tabel := range RegisterTable() {
		errMigrate := db.Debug().AutoMigrate(tabel.Tabel)
		if errMigrate != nil {
			log.Fatal(errMigrate)
		}
	}
}
