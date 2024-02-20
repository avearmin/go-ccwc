package main

import (
	"fmt"
	"log"
	"os"

	"github.com/avearmin/go-ccwc/internal/count"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		return
	}

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

func printCountFromFile(fileName string, f count.CountingFunc) {
	n, err := countFromFile(fileName, f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%8d %s", n, fileName)
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

	fmt.Printf("%8d %8d %8d %s", byteCount, lineCount, wordCount, fileName)
}

func countFromFile(fileName string, f count.CountingFunc) (int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return 0, fmt.Errorf("wc: %s: open: no such file or directory", fileName)
	}

	n := f(file)

	if err = file.Close(); err != nil {
		return 0, err
	}

	return n, nil
}
