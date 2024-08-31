package main

import "fmt"

var ErrTooFewArguments = fmt.Errorf("too few arguments")

type Config struct {
	Repo    string
	Sources []string
	DestDir string
}

func ParseArgs(args []string) (*Config, error) {
	var sources []string
	var repo, destdir string

	if len(args) < 3 {
		return nil, ErrTooFewArguments
	}

	sources = args[1 : len(args)-1]
	repo, destdir = args[0], args[len(args)-1]

	c := &Config{Repo: repo, Sources: sources, DestDir: destdir}
	return c, nil
}
