package command

import (
	"bytes"
	"fmt"
	"github.com/avearmin/go-ccwc/internal/count"
	"io"
	"log"
	"os"
)

func IsValidFlag(flag string) bool {
	flags := map[string]bool{"-c": true, "-l": true, "-w": true, "-m": true}
	ok := flags[flag]
	return ok

}

func HandleStdinFlags(flag string) {
	switch flag {
	case "-c":
		printCountFromStdin(count.Bytes)
	case "-l":
		printCountFromStdin(count.Lines)
	case "-w":
		printCountFromStdin(count.Words)
	case "-m":
		printCountFromStdin(count.Runes)
	}
}

func HandleFileFlags(flag, fileName string) {
	switch flag {
	case "-c":
		printCountFromFile(fileName, count.Bytes)
	case "-l":
		printCountFromFile(fileName, count.Lines)
	case "-w":
		printCountFromFile(fileName, count.Words)
	case "-m":
		printCountFromFile(fileName, count.Runes)
	}
}

func printCountFromFile(fileName string, f count.CountingFunc) {
	n, err := countFromFile(fileName, f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%8d %s\n", n, fileName)
}

func PrintAllCountsFromFile(fileName string) {
	byteCount, err := countFromFile(fileName, count.Bytes)
	if err != nil {
		log.Fatal(err)
	}
	lineCount, err := countFromFile(fileName, count.Lines)
	if err != nil {
		log.Fatal(err)
	}
	wordCount, err := countFromFile(fileName, count.Words)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%8d %8d %8d %s\n", lineCount, wordCount, byteCount, fileName)
}

func countFromFile(fileName string, f count.CountingFunc) (int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return 0, fmt.Errorf("wc: %s: open: no such file or directory\n", fileName)
	}

	n := f(file)

	if err = file.Close(); err != nil {
		return 0, err
	}

	return n, nil
}

func printCountFromStdin(f count.CountingFunc) {
	n := f(os.Stdin)
	fmt.Printf("%8d\n", n)
}

// I dislike having to read the entire input into memory here,
// but currently this is the easiest way to have 3 different
// scanners scan through the data.
func PrintAllCountsFromStdin() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	byteCount := count.Bytes(bytes.NewReader(input))
	lineCount := count.Lines(bytes.NewReader(input))
	wordCount := count.Words(bytes.NewReader(input))

	fmt.Printf("%8d %8d %8d\n", lineCount, wordCount, byteCount)
}
