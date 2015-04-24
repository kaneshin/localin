package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

var commands = []*command{
	cmdGen,
}

func run(args []string) ([]byte, error) {
	if len(args) == 0 {
		usage()
	}
	cmd := exec.Command(args[0], args[1:]...)
	return cmd.Output()
}

func perror(err ...interface{}) {
	fmt.Fprintln(stderr, "ERROR: ", err)
	os.Exit(1)
}

func main() {
	flag.Usage = usage
	flag.Parse()
	if flag.NArg() == 0 {
		usage()
	}

	name := flag.Arg(0)

	if name == "--help" || name == "-h" {
		usage()
	}

	for _, cmd := range commands {
		if cmd.name() == name {
			if err := cmd.run(flag.Args()[1:]); err != nil {
				perror(err)
			}
			os.Exit(0)
		}
	}
	perror("Unknown command ", name)
}
