package count

import (
	"io"
	"strings"
)

type CountingFunc func(io.Reader) (int, error)

func Bytes(r io.Reader) (int, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return 0, err
	}
	return len(data), nil
}

func Lines(r io.Reader) (int, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return 0, err
	}
	lines := strings.Split(string(data), "\n")
	return len(lines), nil
}

func Words(r io.Reader) (int, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return 0, err
	}
	stringData := strings.ReplaceAll(string(data), "\n", " ")
	words := strings.Split(stringData, " ")
	return len(words), nil

}
