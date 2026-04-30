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
