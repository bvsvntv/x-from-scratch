package core

import "errors"

// Reads the length typically the first integer of the string
// until hit by a non-digit byte and returns
// the integer and the delta = length + 2(CRLF)
// TODO: Make it simpler and rad until we get '\r' just like other functions
func readLength(data []byte) (int, int) {
	pos, length := 0, 0

	for pos = range data {
		b := data[pos]
		if !(b >= '0' && b <= '9') {
			return length, pos + 2
		}
		length = length*10 + int(b-'0')
	}

	return 0, 0
}

// Reads a RESP encoded simple string from data and returns
// the string, the delta, and the error
func readSimpleString(data []byte) (string, int, error) {
	// first character '+'
	pos := 1

	for ; data[pos] != '\r'; pos++ {
	}

	return string(data[1:pos]), pos + 2, nil
}

// Reads a RESP encoded error from data and returns
// the error string, the delta, and the error
func readError(data []byte) (string, int, error) {
	return readSimpleString(data)
}

// Reads a RESP encoded integer from data and returns
// the integer value, the delta, and the error
func readInt64(data []byte) (any, int, error) {
	// first character ':'
	pos := 1
	var value int64 = 0

	for ; data[pos] != '\r'; pos++ {
		value = value*10 + int64(data[pos]-'0')
	}

	return value, pos + 2, nil
}

// Reads a RESP encoded string from data and returns
// the string, the delta, and the error
func readBulkString(data []byte) (any, int, error) {
	// first character '$'
	pos := 1

	// reading the length and forwarding the pos by
	// the length of the integer + the first special character
	len, delta := readLength(data[pos:])
	pos += delta

	// reading 'len' bytes a string
	return string(data[pos:(pos + len)]), pos + len + 2, nil
}

func DecodeOne(data []byte) (any, int, error) {
	if len(data) == 0 {
		return nil, 0, errors.New("no data")
	}

	switch data[0] {
	case '+':
		return readSimpleString(data)
	case '-':
		return readError(data)
	case ':':
		return readInt64(data)
	case '$':
		return readBulkString(data)
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
