package cli

import (
	"errors"
	"flag"
	"os"
	"strings"
)

var (
	readFile = os.ReadFile
	getEnv   = os.Getenv
	initArgs = flag.Args
)

var sessionID = flag.String("session", "", "session token from advent of code")

const SESSION_TOKEN = "AOC_SESSION"

type parameters struct {
	SessionID *string
	Url       string
}

func Run() (params *parameters, err error) {
	flag.Parse()
	args := initArgs()
	if len(args) <= 1 || args[0] != "fetch" {
		return params, errors.New("did you want to call fetch?")
	}

	if len(args) < 2 {
		return params, errors.New("please enter a url")
	}

	params = &parameters{
		Url: args[1],
	}

	return params, nil
}

func GrabSessionID() (sessionID string, err error) {
	fileContent, err := readFile("./session")
	if err != nil {
		sessionID = getEnv(SESSION_TOKEN)
	} else {
		sessionID = string(fileContent)
	}

	if len(sessionID) == 0 {
		return sessionID, errors.New("No session id found")
	}

	sessionID = strings.Fields(sessionID)[0]

	return sessionID, nil
}
