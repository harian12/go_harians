package main

import (
	"go_harians/config"
	"go_harians/router"
)

func main() {
	config.COnnectionDB()
	router.AppRouter()
}
