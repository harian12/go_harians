package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func COnnectionDB() *gorm.DB {

	errEnv := godotenv.Load()

	if errEnv != nil {
		panic("gagal load file env")
	}

	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_USER := os.Getenv("DB_USER")
	DB_PASS := os.Getenv("DB_PASS")
	DB_NAME := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME)
	db, errDB := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if errDB != nil {
		panic("gagal koneksi ke Database")
	}

	return db
}
