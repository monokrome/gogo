package main

import (
	"flag"
	"log"
)

type Executer interface {
	Execute([]string) error
}

func main() {
	var executer Executer

	flag.Parse()
	arguments := flag.Args()

	if len(arguments) == 0 {
		flag.PrintDefaults()
		log.Fatalln("A command must be provided.")
	}

	command := arguments[0]

	switch command {
	case "scaffold":
		executer = new(ScaffoldCommand)

	case "serve":
		executer = new(ServeCommand)

	default:
		log.Fatalln("Invalid command:", command)
	}

	executer.Execute(arguments[1:])
}
