package server

import (
	"log"

	"redis/config"
)

func RunSyncTCPServer() {
	log.Printf("starting a synchronous TCP server on %s:%d", config.Host, config.Port)
}
