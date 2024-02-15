package count

import (
	"bufio"
	"strings"
	"testing"
)

func TestBytes(t *testing.T) {
	tests := map[string]struct {
		input *bufio.Scanner
		want  int
	}{
		"12 bytes": {
			bufio.NewScanner(strings.NewReader("Hello world!")),
			12,
		},
		"0 bytes": {
			bufio.NewScanner(strings.NewReader("")),
			0,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if got := Bytes(test.input); test.want != got {
				t.Errorf("FAIL want: %2d got: %2d", test.want, got)
			}
		})
	}
}

func TestLines(t *testing.T) {
	tests := map[string]struct {
		input *bufio.Scanner
		want  int
	}{
		"3 lines": {
			bufio.NewScanner(strings.NewReader("Hello world!\nHello universe!\nHello multiverse!")),
			3,
		},
		"1 lines": {
			bufio.NewScanner(strings.NewReader("")),
			0,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if got := Lines(test.input); test.want != got {
				t.Errorf("FAIL want: %2d got: %2d", test.want, got)
			}
		})
	}
}

func TestWords(t *testing.T) {
	tests := map[string]struct {
		input *bufio.Scanner
		want  int
	}{
		"6 words": {
			bufio.NewScanner(strings.NewReader("Hello world!\nHello universe!\nHello multiverse!")),
			6,
		},
		"0 words": {
			bufio.NewScanner(strings.NewReader("")),
			0,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if got := Words(test.input); test.want != got {
				t.Errorf("FAIL want: %2d got: %2d", test.want, got)
			}
		})
	}
}
