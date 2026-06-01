package core

import (
	"errors"
	"io"
)

func evalPING(args []string, c io.ReadWriter) error {
	var b []byte

	if len(args) >= 2 {
		return errors.New("ERR: wrong number of arguments for 'ping' command")
	}

	if len(args) == 0 {
		b = Encode("PONG", true)
	} else {
		b = Encode(args[0], false)
	}

	_, err := c.Write(b)
	return err
}

func evalSET(args []string, c io.ReadWriter) error {
	if len(args) <= 1 {
		return errors.New("ERR: wrong number of arguments for 'set' command")
	}

	var key, value string
	key, value = args[0], args[1]

	// Put the k key and value in Hash Table
	Put(key, NewObj(value))
	c.Write([]byte("+OK\r\n"))

	return nil
}

func EvalAndRespond(cmd *RedisCmd, c io.ReadWriter) error {
	switch cmd.Cmd {
	case "PING":
		return evalPING(cmd.Args, c)
	case "SET":
		return evalSET(cmd.Args, c)
	default:
		return evalPING(cmd.Args, c)
	}
}
