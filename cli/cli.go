package cli

import (
	"errors"
	"os"
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

func GrabSessionId() (sessionId string, err error) {
	fileContent, err := readFile("./session")
	if err != nil {
		sessionId = getEnv("SESSION")
	} else {
		sessionId = string(fileContent)
	}

	if len(sessionId) == 0 {
		return sessionId, errors.New("No session id found")
	}

	return sessionId, nil
}
