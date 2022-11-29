package fetcher

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

// Tests for Url Validater
func TestURLValidation (t *testing.T) {

	t.Run("", func TestValidateUrl(t *testing.T) {
		t.Parallel()
		url := "https://adventofcode.com/2022/day/1"

		if err := validateURL(url); err != nil {
			t.Error(err)
		}
	})

	t.Run("", func TestErrorIfNotUrl(t *testing.T) {
		t.Parallel()
		url := "123"

		if err := validateURL(url); err == nil {
			t.Errorf("%s should not be a valid url", url)
		}
	})

	t.Run("", func TestErrorIfNotAdventOfCode(t *testing.T) {
		t.Parallel()
		url := "https://google.com"

		if err := validateURL(url); err == nil {
			t.Errorf("%s should return an error", url)
		}

	})

	t.Run("", func TestErrorIfNoPath(t *testing.T) {
		t.Parallel()
		url := "https://adventofcode.com"

		if err := validateURL(url); err == nil {
			t.Errorf("%s should return an error", url)
		}
	})

	t.Run("", func TestErrorIfPathTooShort(t *testing.T) {
		t.Parallel()
		url := "https://adventofcode.com"
		path := "/1/2"

		if err := validateURL(fmt.Sprintf("%s%s", url, path)); err == nil {
			t.Errorf("Should have error because %s is too short", path)
		}
	})

	t.Run("", func TestErrorIfPathTooLong(t *testing.T) {
		t.Parallel()
		url := "https://adventofcode.com"
		path := "/1/2/3/4"

		if err := validateURL(fmt.Sprintf("%s%s", url, path)); err == nil {
			t.Errorf("Should have error because %s is too long", path)
		}
	})

	t.Run("", func TestErrorIfNoYear(t *testing.T) {
		t.Parallel()
		url := "http://adventofcode.com/"
		path := "not/a/year"

		if err := validateURL(fmt.Sprintf("%s%s", url, path)); err == nil {
			t.Errorf("Should have error because %s does not have a year", path)
		}
	})

	t.Run("", func TestErrorOnEarlyYear(t *testing.T) {
		t.Parallel()
		url := "http://adventofcode.com/"
		const year = 2014

		if err := validateURL(fmt.Sprintf("%s%d", url, year)); err == nil {
			t.Errorf("%d should be an invalid year", year)
		}
	})

	t.Run("", func TestErrorOnLateYear(t *testing.T) {
		t.Parallel()
		url := "http://adventofcode.com/"
		const year = 2023

		if err := validateURL(fmt.Sprintf("%s%d", url, year)); err == nil {
			t.Errorf("%d should be an invalid year", year)
		}
	})

	t.Run("", func TestErrorIfDayNotInPath(t *testing.T) {
		t.Parallel()
		url := "http://adventofcode.com/2021"
		notDay := "/not/1"

		if err := validateURL(fmt.Sprintf("%s%s", url, notDay)); err == nil {
			t.Error("Should be an error when no day in url", url)
		}
	})

	t.Run("", func TestErrorIfDayTooLow(t *testing.T) {
		t.Parallel()
		url := "http://adventofcode.com/2021/day/"
		day := 0

		if err := validateURL(fmt.Sprintf("%s%d", url, day)); err == nil {
			t.Errorf("%d should not be a valid day", day)
		}
	})

	t.Run("", func TestErrorIfDayTooHigh(t *testing.T) {
		t.Parallel()
		url := "http://adventofcode.com/2021/day/"
		day := 26

		if err := validateURL(fmt.Sprintf("%s%d", url, day)); err == nil {
			t.Errorf("%d should not be a valid day", day)
		}
	})

}
// Tests for Cookie checker

func TestCookieChecker(t *testing.T) {
	t.Run("", func TestValidCookie(t *testing.T) {
		t.Parallel()
		cookie := http.Cookie{
			Name:  "session",
			Value: "abcdef12345",
		}

		if err := checkCookie(cookie); err != nil {
			t.Error("Should be a valid cookie")
		}
	})

	t.Run("", func TestErrorIfNoSessionCookie(t *testing.T) {
		t.Parallel()
		cookie := http.Cookie{}

		if err := checkCookie(cookie); err == nil {
			t.Error("Should return error if no session cookie")
		}
	})

	t.Run("", func TestErrorIfNoSessionCookieValue(t *testing.T) {
		t.Parallel()
		cookie := http.Cookie{
			Name: "session",
		}

		if err := checkCookie(cookie); err == nil {
			t.Error("Should return error if no value for session cookie")
		}
	})

	t.Run("", func TestErrorIfSessionCookieNotAlphaNumeric(t *testing.T) {
		t.Parallel()
		cookie := http.Cookie{
			Name:  "session",
			Value: "abc123!?.",
		}

		if err := checkCookie(cookie); err == nil {
			t.Error("Should return error if session cookie not alphaumeric")
		}
	})

	t.Run("", func TestErrorIfCookieIsExpired(t *testing.T) {
		t.Parallel()
		cookie := http.Cookie{
			Name:    "session",
			Value:   "abc123",
			Expires: time.Unix(0, 0),
		}

		if err := checkCookie(cookie); err == nil {
			t.Error("Should return error on expired cookie")
		}
	})
}
