package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/avearmin/go-ccwc/internal/count"
)

func main() {
	args := os.Args[1:]

	if len(args) < 2 {
		return
	}

	switch args[0] {
	case "-c":
		countFromFile(args[1], count.Bytes)
	case "-l":
		countFromFile(args[1], count.Lines)
	case "-w":
		countFromFile(args[1], count.Words)
	}
}

func countFromFile(fileName string, f count.CountingFunc) {
	file, err := os.Open(fileName)
	scnr := bufio.NewScanner(file)
	if err != nil {
		log.Fatalf("wc: %s: open: no such file or directory", fileName)
	}
	n := f(scnr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%8d %s", n, fileName)
}
