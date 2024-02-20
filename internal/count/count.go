package count

import (
	"bufio"
	"io"
	"strings"
)

type CountingFunc func(reader io.Reader) int

func Bytes(r io.Reader) int {
	scnr := bufio.NewScanner(r)
	scnr.Split(bufio.ScanBytes)

	var count int
	for scnr.Scan() {
		count++
	}
	return count
}

func Lines(r io.Reader) int {
	scnr := bufio.NewScanner(r)
	scnr.Split(bufio.ScanLines)

	var count int
	for scnr.Scan() {
		count++
	}
	return count
}

func Words(r io.Reader) int {
	scnr := bufio.NewScanner(r)
	scnr.Split(bufio.ScanWords)

	var count int
	for scnr.Scan() {
		word := scnr.Text()
		if strings.TrimSpace(word) == "" {
			continue
		}
		count++
	}
	return count
}

func Runes(r io.Reader) int {
	scnr := bufio.NewScanner(r)
	scnr.Split(bufio.ScanRunes)

	var count int
	for scnr.Scan() {
		count++
	}
	return count
}
