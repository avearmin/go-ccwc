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
		"0 bytes": {
			strings.NewReader(""),
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
		input io.Reader
		want  int
	}{
		"3 lines": {
			strings.NewReader("Hello world!\nHello universe!\nHello multiverse!"),
			3,
		},
		"1 lines": {
			strings.NewReader(""),
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
		input io.Reader
		want  int
	}{
		"6 words": {
			strings.NewReader("Hello world!\nHello universe!\nHello multiverse!"),
			6,
		},
		"0 words": {
			strings.NewReader(""),
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

func TestRunes(t *testing.T) {
	tests := map[string]struct {
		input io.Reader
		want  int
	}{
		"3 lines": {
			strings.NewReader("Hello world!\nHello universe!\nHello multiverse!"),
			46,
		},
		"1 lines": {
			strings.NewReader(""),
			0,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if got := Runes(test.input); test.want != got {
				t.Errorf("FAIL want: %2d got: %2d", test.want, got)
			}
		})
	}
}
