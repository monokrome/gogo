package main

import (
	"flag"
	"log"
)

type Commander interface {
	Execute([]string) error
}

func main() {
	// TODO: How to reference a specific type?
	// Do I **need** to instantiate each one?
	var commands = map[string]Commander{
		"scaffold": new(ScaffoldCommand),
	}

	flag.Parse()
	arguments := flag.Args()

	if len(arguments) == 0 {
		flag.PrintDefaults()
		log.Fatalln("A command must be provided.")
	}

	commandName := arguments[0]
	command, ok := commands[commandName]

	if ok == false {
		log.Fatalln(commandName, "is not a valid command.")
	}

	command.Execute(arguments[1:])
}
