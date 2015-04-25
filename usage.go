package main

import (
	"fmt"
	"os"
)

const usageText = `localin

  Usage:
    localin -h/--help
    localin <command> [arguments...] [options...]

  Commands:
    gen
    lint

  Examples:
    localin gen foo.json bar.json -d localized -f ios android

  General Options:
`

func usage() {
	fmt.Printf(usageText)
	os.Exit(1)
}
