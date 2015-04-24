package main

import (
	"fmt"
	"os"
)

func usage() {
	fmt.Printf(`localin

  Usage:
    localin -h/--help
    localin <command> [arguments...] [options...]

  Commands:
    gen
    lint

  Examples:
    localin gen foo.json bar.json -d localized -f ios android

  General Options:
`)
	os.Exit(0)
}
