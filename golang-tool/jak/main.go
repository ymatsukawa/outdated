package main

import (
	"os"
	"path/filepath"

	"github.com/ymatsukawa/jak/cmd"
	"github.com/ymatsukawa/jak/internal/format"
)

func main() {
	cmdName := getCmdName()
	if err := runCommand(); err != nil {
		handleError(cmdName, err)
	}
}

func getCmdName() string {
	if len(os.Args) > 1 {
		return os.Args[1]
	}
	return filepath.Base(os.Args[0])
}

func runCommand() error {
	return cmd.Execute()
}

func handleError(cmdName string, err error) {
	format.PrintCommandError(cmdName, err)
	os.Exit(1)
}
