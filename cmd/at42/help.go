package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

// HelpOp describes printing help.
type HelpOp struct{
	command string
}

func (h HelpOp) Run(stdout, _ io.Writer) error {
	if h.command != "" {
		return printCommandUsage(stdout, h.command)
	}
	return printUsage(stdout)
}

func printUsage(out io.Writer) error {
	help := `%PROG% is a tool for managing and testing 42 projects source code.

Usage:
	
  %PROG% <command> [arguments]

The commands are:

  clean       remove object files and cached files
  env         print %PROG% environment information
  install     compile and install tests and dependencies
  list        list projects
  test        test projects
  version     print %PROG% version
  help        print this help message

Use "%PROG% help <command>" for more information about a command.`
	help = strings.ReplaceAll(help, "%PROG%", self())

	_, err := fmt.Fprintf(out, "%s\n", help)
	return errors.Wrap(err, "write error")
}

func printCommandUsage(out io.Writer, command string) error {
	help := cmdHelp(command)
	if help == "" {
		return errors.Errorf("unknown command %q", command)
	}

	help = strings.ReplaceAll(help, "%PROG%", self())
	_, err := fmt.Fprintf(out, "%s\n", help)
	return errors.Wrap(err, "write error")
}

func cmdHelp(cmd string) string {
	switch cmd {
	case "test":
		return `usage: %PROG% test [projects]

'%PROG% test' automates testing for 42 projects. It installs tests and run them.`
	case "version":
		return `usage: %PROG% version

If no files are named on the command line, go version prints its own
version information.`
	default:
		return ""
	}
}

func self() string {
	return filepath.Base(os.Args[0])
}