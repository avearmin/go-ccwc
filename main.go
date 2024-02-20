package main

import (
	"github.com/avearmin/go-ccwc/internal/command"
	"os"
)

func main() {
	args := os.Args[1:]
	switch len(args) {
	case 0:
		command.PrintAllCountsFromStdin()
	case 1:
		command.HandleStdinFlags(args[0])
	case 2:
		command.HandleFileFlags(args[0], args[1])
	}
}
