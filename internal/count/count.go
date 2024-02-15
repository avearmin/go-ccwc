package count

import (
	"bufio"
	"strings"
)

type CountingFunc func(*bufio.Scanner) int

func Bytes(scnr *bufio.Scanner) int {
	scnr.Split(bufio.ScanBytes)

	var count int
	for scnr.Scan() {
		count++
	}
	return count
}

func Lines(scnr *bufio.Scanner) int {
	scnr.Split(bufio.ScanLines)

	var count int
	for scnr.Scan() {
		count++
	}
	return count
}

func Words(scnr *bufio.Scanner) int {
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
