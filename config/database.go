package config

import (
	"fmt"
	"go_harians/app/helpers"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func COnnectionDB() *gorm.DB {

	DB_HOST := helpers.GetEnv("DB_HOST", "localhost")
	DB_PORT := helpers.GetEnv("DB_PORT", "3306")
	DB_USER := helpers.GetEnv("DB_USER", "root")
	DB_PASS := helpers.GetEnv("DB_PASS", "sky")
	DB_NAME := helpers.GetEnv("DB_NAME", "harians")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME)
	db, errDB := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if errDB != nil {
		panic("gagal koneksi ke Database")
	}
	//command.NewMigrate(db)
	//seeders.DBSeed(db)
	return db
}
