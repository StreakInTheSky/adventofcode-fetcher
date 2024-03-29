package main

import (
	"errors"
	"fmt"
	"net/http"
	"testing"
	"time"
)

// Tests for Url Validater
const OUTSIDE_ADVENT_DATE = "2023-01-01 00:00:01" // the date to compare for validation
const INSIDE_ADVENT_DATE = "2022-12-01 00:00:01"  // the date to compare for validation

// Get a time struct from the given today to test against
func getNow(t *testing.T, nowString string) time.Time {
	now, err := time.Parse("2006-01-02 15:04:05", nowString)
	if err != nil {
		t.Fatal(err)
	}

	return now
}

func TestValidateUrl(t *testing.T) {
	url := "https://adventofcode.com/2022/day/1"
	now := getNow(t, INSIDE_ADVENT_DATE)

	if err := validateURL(url, now); err != nil {
		t.Error(err)
	}
}

func TestErrorIfNotUrl(t *testing.T) {
	url := "123"
	now := getNow(t, INSIDE_ADVENT_DATE)

	if err := validateURL(url, now); err == nil {
		t.Errorf("%s should not be a valid url", url)
	}
}

func TestErrorIfNotAdventOfCode(t *testing.T) {
	url := "https://google.com"
	now := getNow(t, INSIDE_ADVENT_DATE)

	if err := validateURL(url, now); err == nil {
		t.Errorf("%s should return an error", url)
	}
}

func TestErrorIfNoPath(t *testing.T) {
	url := "https://adventofcode.com"
	now := getNow(t, INSIDE_ADVENT_DATE)

	if err := validateURL(url, now); err == nil {
		t.Errorf("%s should return an error", url)
	}
}

func TestErrorIfPathTooShort(t *testing.T) {
	url := "https://adventofcode.com"
	path := "/1/2"
	now := getNow(t, INSIDE_ADVENT_DATE)

	if err := validateURL(fmt.Sprintf("%s%s", url, path), now); err == nil {
		t.Errorf("Should have error because %s is too short", path)
	}
}

func TestErrorIfPathTooLong(t *testing.T) {
	url := "https://adventofcode.com"
	path := "/1/2/3/4"
	now := getNow(t, INSIDE_ADVENT_DATE)

	if err := validateURL(fmt.Sprintf("%s%s", url, path), now); err == nil {
		t.Errorf("Should have error because %s is too long", path)
	}
}

func TestErrorIfNoYear(t *testing.T) {
	url := "http://adventofcode.com/"
	path := "not/a/year"
	now := getNow(t, INSIDE_ADVENT_DATE)

	if err := validateURL(fmt.Sprintf("%s%s", url, path), now); err == nil {
		t.Errorf("Should have error because %s does not have a year", path)
	}
}

func TestErrorIfDayNotInPath(t *testing.T) {
	url := "http://adventofcode.com/2021"
	notDay := "/not/1"
	now := getNow(t, INSIDE_ADVENT_DATE)

	if err := validateURL(fmt.Sprintf("%s%s", url, notDay), now); err == nil {
		t.Error("Should be an error when no day in url", url)
	}
}

// Tests for Day Validation
func TestErrorIfDayTooLow(t *testing.T) {
	url := "http://adventofcode.com/2021/day/"
	day := 0
	now := getNow(t, INSIDE_ADVENT_DATE)

	if err := validateURL(fmt.Sprintf("%s%d", url, day), now); err == nil {
		t.Errorf("%d should not be a valid day", day)
	}
}

func TestErrorIfDayTooHigh(t *testing.T) {
	url := "http://adventofcode.com/2021/day/"
	day := 26
	now := getNow(t, INSIDE_ADVENT_DATE)

	if err := validateURL(fmt.Sprintf("%s%d", url, day), now); err == nil {
		t.Errorf("%d should not be a valid day", day)
	}
}

func TestErrorIfDayNotYetOpen(t *testing.T) {
	url := "http://adventofcode.com/2021/day/"
	day := 2
	now := getNow(t, INSIDE_ADVENT_DATE)

	if err := validateURL(fmt.Sprintf("%s%d", url, day), now); err == nil {
		t.Errorf("%d should not be a valid day", day)
	}
}

// Tests for year validation
func TestErrorOnEarlyYear(t *testing.T) {
	const year = 2014
	now := getNow(t, INSIDE_ADVENT_DATE)

	if err := validateYear(year, now); err == nil {
		t.Errorf("%d should be an invalid year", year)
	}
}

func TestErrorOnLateYear(t *testing.T) {
	const year = 2024
	now := getNow(t, INSIDE_ADVENT_DATE)

	if err := validateYear(year, now); err == nil {
		t.Errorf("%d should be an invalid year", year)
	}
}

func TestErrorOnCurrentYearEarlyMonth(t *testing.T) {
	const year = 2023
	now := getNow(t, OUTSIDE_ADVENT_DATE)

	if err := validateYear(year, now); err == nil {
		t.Errorf("%d should be an invalid year for the current date, %s", year, OUTSIDE_ADVENT_DATE)
	}
}

// Tests for Cookie checker
func TestValidCookie(t *testing.T) {
	cookie := http.Cookie{
		Name:  "session",
		Value: "abcdef12345",
	}

	if err := checkCookie(cookie); err != nil {
		t.Error("Should be a valid cookie")
	}
}

func TestErrorIfNoSessionCookie(t *testing.T) {
	cookie := http.Cookie{}

	if err := checkCookie(cookie); err == nil {
		t.Error("Should return error if no session cookie")
	}
}

func TestErrorIfNoSessionCookieValue(t *testing.T) {
	cookie := http.Cookie{
		Name: "session",
	}

	if err := checkCookie(cookie); err == nil {
		t.Error("Should return error if no value for session cookie")
	}
}

func TestErrorIfSessionCookieNotAlphaNumeric(t *testing.T) {
	cookie := http.Cookie{
		Name:  "session",
		Value: "abc123!?.",
	}

	if err := checkCookie(cookie); err == nil {
		t.Error("Should return error if session cookie not alphaumeric")
	}
}

func TestErrorIfCookieIsExpired(t *testing.T) {
	cookie := http.Cookie{
		Name:    "session",
		Value:   "abc123",
		Expires: time.Unix(0, 0),
	}

	if err := checkCookie(cookie); err == nil {
		t.Error("Should return error on expired cookie")
	}
}

// Test Cookie Creation
func TestCreatingCookie(t *testing.T) {
	t.Run("Should create session cookie", func(t *testing.T) {
		sessionID := "abc"

		cookie, _ := makeCookie(sessionID)
		if cookie.Value != sessionID {
			t.Errorf("Expected cookie to have value %s, got %s", sessionID, cookie.Value)
		}
	})

	t.Run("Should return error if no sessionID", func(t *testing.T) {
		var sessionID string

		if _, err := makeCookie(sessionID); err == nil {
			t.Error("Expected error when no sessionID")
		}

	})
}

// Tests Fetching Inputs
func TestFetching(t *testing.T) {
	t.Parallel()

	// These tests cannot run in parallel because they share a global variable
	t.Run("Should not return error with successful request", func(t *testing.T) {
		url := "https://adventofcode.com/2021/day/1"
		cookie := http.Cookie{
			Name:  "session",
			Value: "abc123",
		}

		client = &mockClient{res: http.Response{StatusCode: 200}}
		if _, err := fetch(url, cookie); err != nil {
			t.Error(err)
		}
	})

	t.Run("Should return error for invalid url", func(t *testing.T) {
		url := "http://google.com"
		cookie := http.Cookie{
			Name:  "session",
			Value: "abc123",
		}
		client = &mockClient{}

		if _, err := fetch(url, cookie); err == nil {
			t.Errorf("Should return an error with invalid url: %s", url)
		}
	})

	t.Run("Should return error for invalid cookie", func(t *testing.T) {
		url := "https://adventofcode.com/2021/day/1"
		cookie := http.Cookie{
			Name: "invalid",
		}
		client = &mockClient{}

		if _, err := fetch(url, cookie); err == nil {
			t.Errorf("Should return error with invalid cooke: %s", cookie.String())
		}
	})

	t.Run("Should return error if request has failed", func(t *testing.T) {
		url := "https://adventofcode.com/2021/day/1"
		cookie := http.Cookie{
			Name:  "session",
			Value: "abc123",
		}

		expected := http.Response{
			StatusCode: 404,
		}

		client = &mockClient{
			res: expected,
			err: errors.New("There was an error fetching the site"),
		}

		if res, err := fetch(url, cookie); err == nil {
			t.Errorf("Request should return an error with status code %d, but got %d", expected.StatusCode, res.StatusCode)
		}
	})

	t.Run("Should return result", func(t *testing.T) {
		url := "https://adventofcode.com/2021/day/1"
		cookie := http.Cookie{
			Name:  "session",
			Value: "abc123",
		}

		expectedRes := http.Response{
			StatusCode: 200,
		}
		client = &mockClient{
			res: expectedRes,
		}

		res, err := fetch(url, cookie)
		if err != nil {
			t.Errorf("Should not have an error, but got %s", err)
		}

		if res.StatusCode != expectedRes.StatusCode {
			t.Errorf("Expected res.StatusCode to be %d, instead got %d", expectedRes.StatusCode, res.StatusCode)
		}
	})
}

type mockClient struct {
	res http.Response
	err error
}

func (c *mockClient) Do(req *http.Request) (*http.Response, error) {
	return &c.res, c.err
}
