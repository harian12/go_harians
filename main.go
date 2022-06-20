package main

import (
	"flag"
	"go_harians/config"
	"go_harians/router"
)

func main() {

	flag.Parse()
	arg := flag.Arg(0)
	if arg != "" {
		config.Run()
	} else {
		config.Run()
		router.AppRouter()
	}

}
