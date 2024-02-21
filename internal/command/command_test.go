package command

import "testing"

func TestIsValidFlag(t *testing.T) {
	tests := map[string]struct {
		input string
		want  bool
	}{
		"-c valid": {
			"-c",
			true,
		},
		"-l valid": {
			"-l",
			true,
		},
		"-m valid": {
			"-m",
			true,
		},
		"-w valid": {
			"-w",
			true,
		},
		"-o invalid": {
			"-o",
			false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(*testing.T) {
			if got := IsValidFlag(test.input); got != test.want {
				t.Errorf("FAIL want: %5t got: %5t", test.want, got)
			}
		})
	}
}
