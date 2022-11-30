package fetcher

import (
	"errors"
	"fmt"
	"net/http"
	"testing"
	"time"
)

// Tests for Url Validater

func TestValidateUrl(t *testing.T) {
	url := "https://adventofcode.com/2022/day/1"

	if err := validateURL(url); err != nil {
		t.Error(err)
	}
}

func TestErrorIfNotUrl(t *testing.T) {
	url := "123"

	if err := validateURL(url); err == nil {
		t.Errorf("%s should not be a valid url", url)
	}
}

func TestErrorIfNotAdventOfCode(t *testing.T) {
	url := "https://google.com"

	if err := validateURL(url); err == nil {
		t.Errorf("%s should return an error", url)
	}

}

func TestErrorIfNoPath(t *testing.T) {
	url := "https://adventofcode.com"

	if err := validateURL(url); err == nil {
		t.Errorf("%s should return an error", url)
	}
}

func TestErrorIfPathTooShort(t *testing.T) {
	url := "https://adventofcode.com"
	path := "/1/2"

	if err := validateURL(fmt.Sprintf("%s%s", url, path)); err == nil {
		t.Errorf("Should have error because %s is too short", path)
	}
}

func TestErrorIfPathTooLong(t *testing.T) {
	url := "https://adventofcode.com"
	path := "/1/2/3/4"

	if err := validateURL(fmt.Sprintf("%s%s", url, path)); err == nil {
		t.Errorf("Should have error because %s is too long", path)
	}
}

func TestErrorIfNoYear(t *testing.T) {
	url := "http://adventofcode.com/"
	path := "not/a/year"

	if err := validateURL(fmt.Sprintf("%s%s", url, path)); err == nil {
		t.Errorf("Should have error because %s does not have a year", path)
	}
}

func TestErrorOnEarlyYear(t *testing.T) {
	url := "http://adventofcode.com/"
	const year = 2014

	if err := validateURL(fmt.Sprintf("%s%d", url, year)); err == nil {
		t.Errorf("%d should be an invalid year", year)
	}
}

func TestErrorOnLateYear(t *testing.T) {
	url := "http://adventofcode.com/"
	const year = 2023

	if err := validateURL(fmt.Sprintf("%s%d", url, year)); err == nil {
		t.Errorf("%d should be an invalid year", year)
	}
}

func TestErrorIfDayNotInPath(t *testing.T) {
	url := "http://adventofcode.com/2021"
	notDay := "/not/1"

	if err := validateURL(fmt.Sprintf("%s%s", url, notDay)); err == nil {
		t.Error("Should be an error when no day in url", url)
	}
}

func TestErrorIfDayTooLow(t *testing.T) {
	url := "http://adventofcode.com/2021/day/"
	day := 0

	if err := validateURL(fmt.Sprintf("%s%d", url, day)); err == nil {
		t.Errorf("%d should not be a valid day", day)
	}
}

func TestErrorIfDayTooHigh(t *testing.T) {
	url := "http://adventofcode.com/2021/day/"
	day := 26

	if err := validateURL(fmt.Sprintf("%s%d", url, day)); err == nil {
		t.Errorf("%d should not be a valid day", day)
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

// Tests Fetching Inputs
func TestFetching(t *testing.T) {
	t.Parallel()
	t.Run("Should not return error with successful request", func(t *testing.T) {
		t.Parallel()

		url := "https://adventofcode.com/2021/day/1"
		cookie := http.Cookie{
			Name:  "session",
			Value: "abc123",
		}
		if _, err := Fetch(url, cookie); err != nil {
			t.Error(err)
		}
	})

	t.Run("Should return error for invalid url", func(t *testing.T) {
		t.Parallel()

		url := "http://google.com"
		cookie := http.Cookie{
			Name:  "session",
			Value: "abc123",
		}

		if _, err := Fetch(url, cookie); err == nil {
			t.Errorf("Should return an error with invalid url: %s", url)
		}
	})

	t.Run("Should return error for invalid cookie", func(t *testing.T) {
		t.Parallel()

		url := "https://adventofcode.com/2021/day/1"
		cookie := http.Cookie{
			Name: "invalid",
		}

		if _, err := Fetch(url, cookie); err == nil {
			t.Errorf("Should return error with invalid cooke: %s", cookie.String())
		}
	})

	t.Run("Should return error if request has failed", func(t *testing.T) {
		t.Parallel()

		url := "https://adventofcode.com/2021/day/1"
		cookie := http.Cookie{
			Name:  "session",
			Value: "abc123",
		}

		res := http.Response{
			StatusCode: 404,
		}
		Client = &mockClient{
			res: res,
			err: errors.New("There was an error fetching the site"),
		}

		if _, err := Fetch(url, cookie); err == nil {
			t.Error("Response should be an error")
		}
	})

	t.Run("Should return result", func(t *testing.T) {
		t.Parallel()

		url := "https://adventofcode.com/2021/day/1"
		cookie := http.Cookie{
			Name:  "session",
			Value: "abc123",
		}

		expectedRes := http.Response{
			StatusCode: 200,
		}
		Client = &mockClient{
			res: expectedRes,
		}

		res, err := Fetch(url, cookie)
		if err != nil {
			t.Error("Response should be an error")
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
