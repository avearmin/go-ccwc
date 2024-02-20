package main

import (
	"bytes"
	"fmt"
	"github.com/avearmin/go-ccwc/internal/count"
	"io"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		printAllCountsFromStdin()
	} else if len(args) < 2 {
		switch args[0] {
		case "-c":
			printCountFromStdin(count.Bytes)
		case "-l":
			printCountFromStdin(count.Lines)
		case "-w":
			printCountFromStdin(count.Words)
		case "-m":
			printCountFromStdin(count.Runes)
		}
	} else {
		switch args[0] {
		case "-c":
			printCountFromFile(args[1], count.Bytes)
		case "-l":
			printCountFromFile(args[1], count.Lines)
		case "-w":
			printCountFromFile(args[1], count.Words)
		case "-m":
			printCountFromFile(args[1], count.Runes)
		default:
			printAllCountsFromFile(args[0])
		}
	}
}

func printCountFromFile(fileName string, f count.CountingFunc) {
	n, err := countFromFile(fileName, f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%8d %s\n", n, fileName)
}

func printAllCountsFromFile(fileName string) {
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
func printAllCountsFromStdin() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	byteCount := count.Bytes(bytes.NewReader(input))
	lineCount := count.Lines(bytes.NewReader(input))
	wordCount := count.Words(bytes.NewReader(input))

	fmt.Printf("%8d %8d %8d\n", lineCount, wordCount, byteCount)
}
