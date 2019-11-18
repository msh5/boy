package main

import (
	"fmt"

	"github.com/msh5/boy/interface/cmd"
)

// "version", "commit" and "date" variables are overridden by goreleaser.

var (
	version = "dev"
	commit  = "none"    // nolint:gochecknoglobals
	date    = "unknown" // nolint:gochecknoglobals
)

func formatVersionString() string {
	return fmt.Sprintf("%v, commit %v, built at %v", version, commit, date)
}

func main() {
	cmd.SetVersion(formatVersionString())

	cmd.Execute()
}
