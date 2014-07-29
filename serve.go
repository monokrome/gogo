package main

import "fmt"

type ServeCommand struct{}

func (s *ServeCommand) Execute(arguments []string) error {
	fmt.Printf("Starting server.\n")
	return nil
}
