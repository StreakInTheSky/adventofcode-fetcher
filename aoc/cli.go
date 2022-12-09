package main

import (
	"errors"
	"flag"
	"os"
	"regexp"
	"strings"
)

var (
	readFile = os.ReadFile
	getEnv   = os.Getenv
	initArgs = flag.Args
)

var sessionFlag = flag.String("session", "./session", "session token from advent of code")

const SESSION_TOKEN = "AOC_SESSION"

func run() (url, sessionParam string, err error) {
	flag.Parse()
	args := initArgs()
	if len(args) < 1 || args[0] != "fetch" {
		return url, sessionParam, errors.New("Did you want to call \"fetch\"?")
	}

	if len(args) < 2 {
		return url, sessionParam, errors.New("Please enter a url")
	}

	return args[1], *sessionFlag, nil
}

func isPath(input string) bool {
	return regexp.MustCompile("^[./]").MatchString(input)
}

func grabSessionID(sessionParam string) (sessionID string, err error) {
	err = errors.New("No session id found")

	if isPath(sessionParam) {
		fileContent, err := readFile(sessionParam)
		if err != nil {
			return sessionID, err
		}

		sessionID = strings.Fields(string(fileContent))[0]
	} else {
		sessionID = sessionParam
	}

	if len(sessionID) == 0 {
		return sessionID, err
	}

	return sessionID, nil
}
