// Package fetcher grabs inputs for the user from Advent of Code.
// Validates url inputs for valid Advent of Code urls
// Makes sure there is a cookie for the request
package fetcher

import (
	"errors"
	"fmt"
	"log"
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

func MakeCookie(sessionId string) (cookie http.Cookie, err error) {
	if len(sessionId) == 0 {
		return cookie, errors.New("sessionId must have a value")
	}

	cookie = http.Cookie{
		Name:  "session",
		Value: sessionId,
	}
	return cookie, err
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var Client HTTPClient

func init() {
	Client = &http.Client{}
}

func Fetch(url string, cookie http.Cookie) (res *http.Response, err error) {
	if err = validateURL(url); err != nil {
		return res, err
	}

	if err = checkCookie(cookie); err != nil {
		return res, err
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return res, err
	}

	req.AddCookie(&cookie)
	res, err = Client.Do(req)
	if err != nil {
		return res, err
	}

	log.Print(res.StatusCode)
	if res.StatusCode >= 400 {
		return res, errors.New("There was an error making the request")
	}

	return res, err
}
