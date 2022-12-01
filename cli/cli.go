package cli

import "errors"

func ParseArgs(args []string) (string, error) {
	var url string
	if len(args) == 1 || args[1] != "fetch" {
		return url, errors.New("not enough arguments")
	}

	if len(args) < 3 {
		return url, errors.New("please enter a url")
	}

	return args[2], nil
}
