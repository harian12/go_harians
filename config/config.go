package config

import (
	"flag"
	"github.com/joho/godotenv"
	"go_harians/command"
	"gorm.io/gorm"
	"log"
	"os"
)

type Server struct {
	DB *gorm.DB
}

type DBConfig struct {
	DB_HOST string
	DB_PORT string
	DB_USER string
	DB_PASS string
	DB_NAME string
}

func LoadEnvDB() DBConfig {
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("Error load .env File")
	}
	var dbConfig = DBConfig{}

	dbConfig.DB_HOST = GetEnv("DB_HOST", "localhost")
	dbConfig.DB_PORT = GetEnv("DB_PORT", "3306")
	dbConfig.DB_USER = GetEnv("DB_USER", "root")
	dbConfig.DB_PASS = GetEnv("DB_PASS", "sky")
	dbConfig.DB_NAME = GetEnv("DB_NAME", "harians")

	return dbConfig
}

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func Run() {
	dbConfig := LoadEnvDB()
	connectDB := COnnectionDB(dbConfig)

	arg := flag.Arg(0)
	if arg != "" {
		command.InitCommands(connectDB)
	}
}
