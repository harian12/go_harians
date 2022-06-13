package main

import (
	"go_harians/command"
	"go_harians/config"
	"go_harians/router"
)

func main() {
	config.COnnectionDB()
	command.Arguments()
	router.AppRouter()
}
