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
