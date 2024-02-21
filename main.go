package main

import (
	"github.com/avearmin/go-ccwc/internal/command"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		command.PrintAllCountsFromStdin()
	} else if len(args) == 1 && command.IsValidFlag(args[0]) {
		command.HandleStdinFlags(args[0])
	} else if len(args) == 2 && command.IsValidFlag(args[0]) {
		command.HandleFileFlags(args[0], args[1])
	} else {
		log.Printf("\nwc: illegal option -- %s\nusage: wc [-clmw] [file ...]\n", args[0])
	}
}
