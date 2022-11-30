// Package fetcher grabs inputs for the user from Advent of Code.
// Validates url inputs for valid Advent of Code urls
// Makes sure there is a cookie for the request
package fetcher

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const firstYear = 2015
const currentYear = 2022

func validateURL(inputURL string) error {
	parsedURL, err := url.Parse(inputURL)

	if err != nil {
		return err
	}

	if parsedURL.Host != "adventofcode.com" {
		return fmt.Errorf("%s is not a valid advent of code url", inputURL)
	}

	parsedPath := strings.Split(parsedURL.Path, "/")

	if len(parsedPath) < 4 {
		return errors.New("Url did not include a day")
	}

	year, err := strconv.Atoi(parsedPath[1])
	if err != nil {
		return err
	}
	if year < firstYear || year > currentYear {
		return fmt.Errorf("Invalid year: %d", year)
	}

	if parsedPath[2] != "day" {
		return errors.New("Url does not include day")
	}

	day, err := strconv.Atoi(parsedPath[3])
	if err != nil {
		return err
	}
	if day < 1 || day > 25 {
		return fmt.Errorf("%d is not a valid day", day)
	}

	return nil
}

func checkCookie(cookie http.Cookie) error {
	if cookie.Name != "session" || cookie.Value == "" {
		return errors.New("No session cookie")
	}

	if !regexp.MustCompile(`^[a-zA-Z0-9]*$`).MatchString(cookie.Value) {
		return errors.New("Not a valid session cookie")
	}

	if !cookie.Expires.IsZero() && cookie.Expires.Before(time.Now()) {
		return errors.New("Expired session cookie")
	}

	return nil
}

func Fetch(url string, cookie http.Cookie) error {
	if err := validateURL(url); err != nil {
		return err
	}

	if err := checkCookie(cookie); err != nil {
		return err
	}
	return nil
}
