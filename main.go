// Grabs input from Advent Of Code Events
package main

import (
	"fmt"
	"io"
	"os"

	"github.com/streakinthesky/adventofcode-fetcher/cli"
	"github.com/streakinthesky/adventofcode-fetcher/fetcher"
)

func main() {
	url, err := cli.ParseArgs(os.Args)
	if err != nil {
		handleError(err, 1)
	}

	sessionID, err := cli.GrabSessionID()
	if err != nil {
		handleError(err, 1)
	}

	cookie, err := fetcher.MakeCookie(sessionID)
	if err != nil {
		handleError(err, 1)
	}

	res, err := fetcher.Fetch(url, cookie)
	if err != nil {
		handleError(err, 1)
	}
	defer res.Body.Close()
}

func handleError(err error, exitCode int) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(exitCode)
}

func handleOutput(body io.Reader, target *os.File) (err error) {
	if _, err := io.Copy(target, body); err != nil {
		return err
	}
	return err
}
