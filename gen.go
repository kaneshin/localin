package main

import ()

var cmdGen = &command{
	usage: "clone <repository> [options...]",
}

func init() {
	cmdGen.run = gen
}

func gen(args []string) error {
	if len(args) == 0 {
		return cmdGen.err()
	}
	var files []string
	// var destination string
	// var format []string
	for _, v := range args {
		if v == "-d" {
			// destination = args[i+1]
			continue
		} else if v == "-f" {
			continue
		}
		files = append(files, v)
	}
	return nil
}
