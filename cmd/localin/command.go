package main

import (
	"errors"
	"strings"
)

type command struct {
	usage string
	run   func([]string) error
	// options map[string]func(interface{})
}

func (c *command) name() string {
	if f := strings.Fields(c.usage); len(f) > 0 {
		return f[0]
	}
	return ""
}

func (c *command) err() error {
	return errors.New(`While executing localin ` + c.name() + `
Usage: ` + c.usage)
}
