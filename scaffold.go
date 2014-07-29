package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/VictorLowther/go-git/git"
)

type ScaffoldCommand struct{}

func (s *ScaffoldCommand) Execute(arguments []string) error {
	if len(arguments) == 0 {
		log.Fatalln("Must provide source repository.")
	}

	if len(arguments) == 1 {
		log.Fatalln("Must provide project name.")
	}

	source := arguments[0]
	destination := arguments[1]

	cacheRoot := os.Getenv("HOME")

	if cacheRoot != "" {
		cacheRoot = filepath.Join(cacheRoot, ".config", "gogo")
	} else {
		cacheRoot = filepath.Join(".gogo")
	}

	cachePath := filepath.Join(cacheRoot, source)

	_, err := os.Stat(cachePath)

	if err != nil && os.IsNotExist(err) {
		_, err := git.Clone(source, cachePath)

		if err != nil {
			log.Fatalln(err)
		}
	} else {
		var remotes []string

		repo, err := git.Open(cachePath)

		if err != nil {
			log.Fatalln(err)
		}

		remoteMap := repo.Remotes()

		for k, _ := range remoteMap {
			remotes = append(remotes, k)
		}

		repo.Fetch(remotes)
	}

	command := exec.Command("cp", "-R", cachePath, destination)

	command.Start()

	if err := command.Wait(); err != nil {
		log.Fatalln(err)
	}

	repo, err := git.Open(destination)

	if err != nil {
		log.Fatalln(err)
	}

	if err := repo.ZapRemote("origin"); err != nil {
		log.Fatalln(err)
	}

	if err := repo.AddRemote("base", source); err != nil {
		log.Fatalln(err)
	}

	return nil
}
