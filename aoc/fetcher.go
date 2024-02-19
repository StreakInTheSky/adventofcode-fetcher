package main

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

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var client httpClient = &http.Client{}

// / validates the input URL based on the ETC/UTC-5 date as that is when puzzles are unlocked
func validateURL(inputURL string, now time.Time) error {
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

	if err := validateYear(year, now); err != nil {
		return err
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

	currDay := now.Day()
	if day > currDay {
		return fmt.Errorf("%d is not yet open", day)
	}

	return nil
}

func validateYear(year int, now time.Time) error {
	currentYear := now.Year()
	if year < firstYear || year > currentYear {
		return fmt.Errorf("Invalid year: %d", year)
	}
	currentMonth := now.Month()
	if year == currentYear && currentMonth != time.December {
		return errors.New("It is not December yet")
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

// MakeCookie makes a session cookie from a sessionID
func makeCookie(sessionID string) (cookie http.Cookie, err error) {
	if len(sessionID) == 0 {
		return cookie, errors.New("sessionId must have a value")
	}

	cookie = http.Cookie{
		Name:  "session",
		Value: sessionID,
	}
	return cookie, err
}

// Fetch fetches input for advent of code url and a user's session cookie
func fetch(url string, cookie http.Cookie) (res *http.Response, err error) {
	est, err := time.LoadLocation("America/New_York")
	if err != nil {
		return res, err
	}

	today := time.Now().In(est)
	if err = validateURL(url, today); err != nil {
		return res, err
	}

	if err = checkCookie(cookie); err != nil {
		return res, err
	}

	req, err := http.NewRequest("GET", fmt.Sprint(url, "/input"), nil)
	if err != nil {
		return res, err
	}

	req.AddCookie(&cookie)
	res, err = client.Do(req)
	if err != nil {
		return res, err
	}

	return res, err
}
