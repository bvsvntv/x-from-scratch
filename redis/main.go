package main

import (
	"flag"
	"log"

	"redis/config"
	"redis/server"
)

func setupFlags() {
	flag.StringVar(&config.Host, "host", "0.0.0.0", "host for the redis server")
	flag.IntVar(&config.Port, "port", 2369, "port for the redis server")
	flag.Parse()
}

func main() {
	setupFlags()
	log.Println("starting redis server...")
	server.RunAsyncTCPServer()
}
