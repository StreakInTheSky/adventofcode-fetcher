package cli

import (
	"errors"
	"os"
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

func checkSessionId(sessionId string) (string, error) {
	if len(sessionId) == 0 {
		return sessionId, errors.New("No session id found")
	}

	return sessionId, nil
}

func GrabSessionId() (sessionId string, err error) {
	fileContent, err := os.ReadFile("./session")
	if err != nil {
		sessionId = os.Getenv("SESSION")
	} else {
		sessionId = string(fileContent)
	}

	return checkSessionId(sessionId)
}
