package main

import (
	"fmt"
	"io"

	"github.com/pkg/errors"
)

var (
	version = "v0.0.1"
)

// VersionOps describes printing version string.
type VersionOp struct{}

func (v VersionOp) Run(stdout, _ io.Writer) error {
	_, err := fmt.Fprintf(stdout, "%s\n", version)
	return errors.Wrap(err, "write error")
}
