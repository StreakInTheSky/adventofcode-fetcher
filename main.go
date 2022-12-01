package main

import (
	"fmt"
	"os"

	"github.com/streakinthesky/adventofcode-fetcher/cli"
)

func main() {
	url, err := cli.ParseArgs(os.Args)
	if err != nil {
		handleError(err, 1)
	}

	println(url)
	return
}

func handleError(err error, exitCode int) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(exitCode)
}
