package core

import (
	"testing"
)

func TestSimpleStringDecode(t *testing.T) {
	cases := map[string]string{
		"+OK\r\n": "OK",
	}

	for k, v := range cases {
		value, _ := Decode([]byte(k))
		if v != value {
			t.Fail()
		}
	}
}

func TestErrorDecode(t *testing.T) {
	cases := map[string]string{
		"-Error message\r\n": "Error message",
	}

	for k, v := range cases {
		value, _ := Decode([]byte(k))
		if v != value {
			t.Fail()
		}
	}
}

func TestInt64Decode(t *testing.T) {
	cases := map[string]int64{
		":0\r\n":    0,
		":1000\r\n": 1000,
	}

	for k, v := range cases {
		value, _ := Decode([]byte(k))
		if v != value {
			t.Fail()
		}
	}
}

func TestBulkStringDecode(t *testing.T) {
	cases := map[string]string{
		"$5\r\nhello\r\n": "hello",
		"$0\r\n\r\n":      "",
	}

	for k, v := range cases {
		value, _ := Decode([]byte(k))
		if v != value {
			t.Fail()
		}
	}
}
