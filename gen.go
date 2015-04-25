package main

import (
	"fmt"
)

var cmdGen = &command{
	usage: "gen <repository> [options...]",
}

func init() {
	cmdGen.run = gen
}

func gen(args []string) error {
	if len(args) == 0 {
		return cmdGen.err()
	}
	var files []string = args[:]
	fmt.Println(files)
	return nil
}
