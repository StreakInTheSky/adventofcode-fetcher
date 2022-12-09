// Grabs input from Advent Of Code Events
package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

var (
	createFile = os.Create
)

func main() {
	url, sessionFlag, err := run()
	if err != nil {
		handleError(err, 2)
	}

	sessionID, err := grabSessionID(sessionFlag)
	if err != nil {
		handleError(err, 2)
	}

	cookie, err := makeCookie(sessionID)
	if err != nil {
		handleError(err, 1)
	}

	res, err := fetch(url, cookie)
	if err != nil {
		handleError(err, 18)
	}
	defer res.Body.Close()

	file, err := createOutputFile("inputs.txt")
	if err != nil {
		handleError(err, 18)
	}

	if err := handleOutput(res.Body, file); err != nil {
		handleError(err, 18)
	}
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

func createOutputFile(name string) (file *os.File, err error) {
	exists, err := checkFileExist(name)
	if err != nil {
		return file, err
	}
	if exists {
		return file, errors.New("inputs.txt already exists in the directory")
	}
	return createFile(name)
}

func checkFileExist(name string) (bool, error) {
	if _, err := os.Stat(name); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
