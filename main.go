package main

import (
	"flag"
	"log"

	"github.com/saanny/goddi-db/config"
	"github.com/saanny/goddi-db/server"
)

func setupFlags() {
	flag.StringVar(&config.Host, "host", "0.0.0.0", "host for the dice server")
	flag.IntVar(&config.Port, "port", 7379, "port for the dice server")
	flag.Parse()
}

func main() {
	setupFlags()
	log.Println("Welcome to goddi-db POW POW ")
	server.RunSyncTCPServer()
}
