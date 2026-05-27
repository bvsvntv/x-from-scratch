package server

import (
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"strings"

	"redis/config"
	"redis/core"
)

func readCommand(c io.ReadWriter) (*core.RedisCmd, error) {
	var buf []byte = make([]byte, 512)
	n, err := c.Read(buf[:])
	if err != nil {
		return nil, err
	}

	tokens, err := core.DecodeArrayString(buf[:n])
	if err != nil {
		return nil, err
	}

	return &core.RedisCmd{
		Cmd:  strings.ToUpper(tokens[0]),
		Args: tokens[1:],
	}, nil
}

func respondError(err error, c io.ReadWriter) {
	fmt.Fprintf(c, "-%s\r\n", err)
}

func respond(cmd *core.RedisCmd, c io.ReadWriter) {
	err := core.EvalAndRespond(cmd, c)
	if err != nil {
		respondError(err, c)
	}
}

func RunSyncTCPServer() {
	log.Printf("starting a synchronous TCP server on %s:%d", config.Host, config.Port)

	var con_clients int = 0

	// listening to the configures host:port
	lsnr, err := net.Listen("tcp", config.Host+":"+strconv.Itoa(config.Port))
	if err != nil {
		log.Println(err)
		return
	}

	for {
		// Blocking call: waiting for the new client to connect
		c, err := lsnr.Accept()
		if err != nil {
			log.Println(err)
		}

		// increment the number concurrent clients
		con_clients += 1

		for {
			// Over the socket, continuously read the command print it out
			cmd, err := readCommand(c)
			if err != nil {
				c.Close()
				con_clients -= 1
				if err == io.EOF {
					break
				}
			}

			respond(cmd, c)
		}
	}
}
