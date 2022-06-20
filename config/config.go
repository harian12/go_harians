package config

import (
	"flag"
	"github.com/joho/godotenv"
	"go_harians/command"
	"log"
)

func Run() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("Error load .env File")
	}
	connectDB := COnnectionDB()

	arg := flag.Arg(0)
	if arg != "" {
		command.InitCommands(connectDB)
	}
}
