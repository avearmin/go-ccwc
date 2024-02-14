package count

import (
	"io"
)

type CountingFunc func(io.Reader) (int, error)

func Bytes(r io.Reader) (int, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return 0, err
	}
	return len(data), nil
}
