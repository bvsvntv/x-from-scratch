package core

import "errors"

// Reads a RESP encoded simple string from data and returns
// the string, the delta, and the error
func readSimpleString(data []byte) (string, int, error) {
	// first character '+'
	pos := 1

	for ; data[pos] != '\r'; pos++ {
	}

	return string(data[1:pos]), pos + 2, nil
}

func DecodeOne(data []byte) (any, int, error) {
	if len(data) == 0 {
		return nil, 0, errors.New("no data")
	}

	switch data[0] {
	case '+':
		return readSimpleString(data)
	case '-':
		// read error
	case ':':
		// read int
	case '$':
		// read bulk string
	case '*':
		// read array
	}

	return nil, 0, nil
}

func Decode(data []byte) (any, error) {
	if len(data) == 0 {
		return nil, errors.New("no data")
	}

	value, _, err := DecodeOne(data)

	return value, err
}
