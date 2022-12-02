// Grabs input from Advent Of Code Events
package main

import (
	"fmt"
	"os"

	"github.com/streakinthesky/adventofcode-fetcher/cli"
	"github.com/streakinthesky/adventofcode-fetcher/fetcher"
)

func main() {
	url, err := cli.ParseArgs(os.Args)
	if err != nil {
		handleError(err, 1)
	}

	sessionID, err := cli.GrabSessionId()
	if err != nil {
		handleError(err, 1)
	}

	cookie, err := fetcher.MakeCookie(sessionID)
	if err != nil {
		handleError(err, 1)
	}

	_, err = fetcher.Fetch(url, cookie)
	if err != nil {
		handleError(err, 1)
	}
}

func handleError(err error, exitCode int) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(exitCode)
}
