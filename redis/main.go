package main

import (
	"flag"
	"log"

	"redis/config"
	"redis/server"
)

func main() {
	flag.StringVar(&config.Host, "host", "0.0.0.0", "host for the redis server")
	flag.IntVar(&config.Port, "port", 2369, "port for the redis server")
	flag.Parse()

	log.Println("starting redis server...")
	server.RunSyncTCPServer()
}
