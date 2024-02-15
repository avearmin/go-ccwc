package count

import (
	"io"
	"strings"
	"testing"
)

func TestBytes(t *testing.T) {
	tests := map[string]struct {
		input io.Reader
		want  int
	}{
		"12 bytes": {
			strings.NewReader("Hello world!"),
			12,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if got, _ := Bytes(test.input); test.want != got {
				t.Errorf("FAIL want: %2d got: %2d", test.want, got)
			}
		})
	}
}

func TestLines(t *testing.T) {
	tests := map[string]struct {
		input io.Reader
		want  int
	}{
		"3 lines": {
			strings.NewReader("Hello world!\nHello universe!\nHello multiverse!"),
			3,
		},
		"1 lines": {
			strings.NewReader(""),
			1,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if got, _ := Lines(test.input); test.want != got {
				t.Errorf("FAIL want: %2d got: %2d", test.want, got)
			}
		})
	}
}
