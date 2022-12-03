package cli

import (
	"errors"
	"os"
	"strings"
)

var (
	readFile = os.ReadFile
	getEnv   = os.Getenv
)

func ParseArgs(args []string) (string, error) {
	var url string
	if len(args) <= 1 || args[1] != "fetch" {
		return url, errors.New("did you want to call fetch?")
	}

	if len(args) < 3 {
		return url, errors.New("please enter a url")
	}

	return args[2], nil
}

func GrabSessionID() (sessionID string, err error) {
	fileContent, err := readFile("./session")
	if err != nil {
		sessionID = getEnv("SESSION")
	} else {
		sessionID = string(fileContent)
	}

	if len(sessionID) == 0 {
		return sessionID, errors.New("No session id found")
	}

	sessionID = strings.Fields(sessionID)[0]

	return sessionID, nil
}
