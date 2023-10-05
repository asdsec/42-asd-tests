package main

import (
	"fmt"
	"io"
	"os"

	"github.com/asdsec/at42/internal/env"
	"github.com/asdsec/at42/internal/printer"
	"github.com/fatih/color"
)

type Operation interface {
	Run(stdout, stderr io.Writer) error
}

func main() {
	op := parseArgs(os.Args[1:])
	if err := op.Run(color.Output, color.Error); err != nil {
		printer.Error(color.Error, err.Error())

		if _, ok := os.LookupEnv(env.EnvDebug); ok {
			// print stack trace in verbose mode
			fmt.Fprintf(color.Error, "[DEBUG] error: %+v\n", err)
		}
		defer os.Exit(1)
	}
}