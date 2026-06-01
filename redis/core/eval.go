package core

import (
	"errors"
	"io"
	"strconv"
	"time"
)

var RESP_NIL []byte = []byte("$-1\r\n")

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
	var exDurationMs int64 = -1

	key, value = args[0], args[1]

	for i := 2; i < len(args); i++ {
		switch args[i] {
		case "EX", "ex":
			i++
			if i == len(args) {
				return errors.New("ERR: syntax error")
			}

			exDurationSec, err := strconv.ParseInt(args[3], 10, 64)
			if err != nil {
				return errors.New("ERR: value is not an integer of out of range")
			}
			exDurationMs = exDurationSec * 1000
		default:
			return errors.New("ERR: syntax error")
		}
	}

	// Put the k key and value in Hash Table
	Put(key, NewObj(value, exDurationMs))
	c.Write([]byte("+OK\r\n"))

	return nil
}

func evalGET(args []string, c io.ReadWriter) error {
	if len(args) != 1 {
		return errors.New("ERR: wrong number of arguments for 'get' command")
	}

	var key string = args[0]

	// Get the key from the Hash Table
	obj := Get(key)

	// If key does not exist, return RESP encoded nil
	if obj == nil {
		c.Write(RESP_NIL)
		return nil
	}

	// If key already expired then return nil
	if obj.ExpiresAt != -1 && obj.ExpiresAt <= time.Now().UnixMilli() {
		c.Write(RESP_NIL)
		return nil
	}

	// Return value
	c.Write(Encode(obj.Value, false))
	return nil
}

func EvalAndRespond(cmd *RedisCmd, c io.ReadWriter) error {
	switch cmd.Cmd {
	case "PING":
		return evalPING(cmd.Args, c)
	case "SET":
		return evalSET(cmd.Args, c)
	case "GET":
		return evalGET(cmd.Args, c)
	default:
		return evalPING(cmd.Args, c)
	}
}
