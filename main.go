package main

import (
	"flag"
	"os"
	"fmt"
	"./command"
	"log"
	"strings"
	"io/ioutil"
)

func main() {
	opts := parseFlags()
	log.Print(opts)

	getHosts(opts["hostStr"].(string), opts["hostFile"].(string))
}

func getHosts(hostStr string, hostFile string) (hosts []string) {
	hosts = []string{}
	if len(strings.TrimSpace(hostStr)) > 0 {
		hostsFromStr := strings.Split(hostStr, ",")
		hosts = append(hosts, hostsFromStr...)
	}
	if hostFile != "" {
		content, err := ioutil.ReadFile(hostFile)
		if err != nil {
			log.Fatal("could not read hosts from file")
		}
		if len(content) > 0 {
			hostsFromFile := strings.Split(string(content), "\n")
			for _, oneLine := range hostsFromFile {
				if len(strings.TrimSpace(oneLine)) <= 0 {
					continue
				}
				if oneLine[0] == '#' {
					continue
				}
				hosts = append(hosts, oneLine)
			}
		}
	}
	log.Printf("found %d hosts.", len(hosts))

	return hosts
}

func parseFlags() (opts map[string]interface{}) {
	options := map[string]interface{}{}

	runCommand := flag.NewFlagSet("run", flag.ExitOnError)
	hostStr := runCommand.String("h", "", "hosts seperated by ,")
	hostFile := runCommand.String("f", "", "/path/to/host/file")

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

		pwd, err := os.Getwd()
		if err != nil {
			log.Fatal("fail to get current working directory.")
		}
		tmplFile := strings.Replace(commandModule, ".", string(os.PathSeparator), 1)
		tmplFile = strings.Join([]string{pwd, tmplFile}, string(os.PathSeparator))
		tmplFile = strings.Join([]string{tmplFile, ".json"}, "")
		options["tmplFile"] = tmplFile
		log.Printf("using command file: %s", tmplFile)
		if len(os.Args) == 3 {
			cmdTmpl, err := command.Loadf(tmplFile)
			if err != nil {
				log.Fatal("error: ", err)
			}
			cmdTmpl.PrintUsage()
			os.Exit(1)
		} else {
			varIndex := 0
			for ;varIndex < len(os.Args); varIndex++ {
				varArg := os.Args[varIndex]
				if varArg[0:2] == "--" {
					break
				}
			}
			
			runCommand.Parse(os.Args[3:varIndex])
			if len(*hostStr) <= 0 && len(*hostFile) <= 0 {
				log.Fatal("-h or -f must be provided.")
			}

			options["hostStr"] = *hostStr
			options["hostFile"] = *hostFile
			varMap := map[string]string{}
			if varIndex < len(os.Args) {
				otherArgs := os.Args[varIndex:]
				for index := 0; index < len(otherArgs); index += 1 {
					oKey := otherArgs[index]
					if oKey[0:1] == "--" {
						oKey = oKey[2:]
					} else {
						//continue
					}
					oVal := otherArgs[index+1]
					varMap[oKey] = oVal
					index += 1
				}
			}
			options["@vars"] = varMap
		}
	}

	return options
}
