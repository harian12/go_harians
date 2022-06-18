package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func COnnectionDB(dbConfig DBConfig) *gorm.DB {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbConfig.DB_USER, dbConfig.DB_PASS, dbConfig.DB_HOST, dbConfig.DB_PORT, dbConfig.DB_NAME)
	db, errDB := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if errDB != nil {
		panic("gagal koneksi ke Database")
	}
	//command.NewMigrate(db)
	//seeders.DBSeed(db)
	return db

}
