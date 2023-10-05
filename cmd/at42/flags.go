package main

import (
	"fmt"
	"io"
)

type UnsupportedOp struct{ Err error }

func (op UnsupportedOp) Run(_, _ io.Writer) error {
	return op.Err
}

// parseArgs parses the command line arguments and returns the corresponding
func parseArgs(argv []string) Operation {
	if len(argv) == 0 {
		return HiOp{}
	}
	if len(argv) == 1 {
		v := argv[0]
		if v == "help" || v == "--help" || v == "-h" {
			return HelpOp{}
		}
		if v == "version" {
			return VersionOp{}
		}
		return UnsupportedOp{Err: fmt.Errorf("unknown command %q", v)}
	}
	if len(argv) == 2 {
		v := argv[0]
		flag := argv[1]
		if v == "help" || v == "--help" || v == "-h" {
			return HelpOp{command: flag}
		}
		if v == "test" {
			return TestOp{Project: flag}
		}
		return UnsupportedOp{Err: fmt.Errorf("unknown command %q", v)}
	}
	return UnsupportedOp{Err: fmt.Errorf("too many arguments")}
}
