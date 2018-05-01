package main

import (
	"flag"
	"os"
	"fmt"
	"./command"
	"log"
	"strings"
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
		commandModule := os.Args[2]
		if len(os.Args) == 3 {
			pwd, err := os.Getwd()
			if err != nil {
				log.Fatal("fail to get current working directory.")
			}

			tmplFile := strings.Replace(commandModule, ".", string(os.PathSeparator), 1)
			tmplFile = strings.Join([]string{pwd, tmplFile}, string(os.PathSeparator))
			tmplFile = strings.Join([]string{tmplFile, ".json"}, "")
			log.Printf("using command file: %s", tmplFile)
			cmdTmpl, err := command.Loadf(tmplFile)
			if err != nil {
				log.Fatal("error: ", err)
			}
			cmdTmpl.PrintUsage()
		} else {
			runCommand.Parse(os.Args[3:])
		}
	}
}
