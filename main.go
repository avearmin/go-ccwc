package main

import (
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
	}
}

func countFromFile(fileName string, f count.CountingFunc) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("wc: %s: open: no such file or directory", fileName)
	}
	n, err := f(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%8d %s", n, fileName)
}
