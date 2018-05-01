package main

import (
	"flag"
	"os"
	"fmt"
)

func main() {
	parseFlags()
}

func parseFlags() {
	runCommand := flag.NewFlagSet("run", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Printf("Usage: gnuc help|run|create")
		os.Exit(1)
	}

	if os.Args[1] == "help" {
		if len(os.Args) == 2 {
			flag.PrintDefaults()
			os.Exit(0)
		}

		subcommand := os.Args[2]
		if subcommand == "run" {
			runCommand.PrintDefaults()
		}
		os.Exit(0)
	} else if os.Args[1] == "run" {
		if len(os.Args) == 2 {
			runCommand.PrintDefaults()
			os.Exit(0)
		}

	}
}
