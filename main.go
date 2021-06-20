package main

import (
	"fmt"
	"os"

	"go.thunderatz.org/tleague/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
