package main

import (
	"fmt"
	"io"

	"github.com/asdsec/at42/internal/printer"
)

// TestOp indicates the intention to test a project.
type TestOp struct{
	Project string // Project to test
}

func (t TestOp) Run(stdout, _ io.Writer) error {
	ps := fetchProjects()
	for _,p := range ps {
		if  p != t.Project {
			return  fmt.Errorf("project %s not found", t.Project)
		}
		printer.ActiveItemColor.Println(p)
	}
	return nil
}

// fetchProjects returns the list of projects available for testing.
func fetchProjects() []string {
	return []string{
		"libft",
		"printf",
		"get_next_line",
	}
}
