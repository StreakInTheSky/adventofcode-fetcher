package fetcher

import (
	"fmt"
	"testing"
)

func TestValidateUrl(t *testing.T) {
	url := "https://adventofcode.com/2022/day/1"

	if err := validate_url(url); err != nil {
		t.Error(err)
	}
}

func TestErrorIfNotUrl(t *testing.T) {
	url := "123"

	if err := validate_url(url); err == nil {
		t.Errorf("%s should not be a valid url", url)
	}
}

func TestErrorIfNotAdventOfCode(t *testing.T) {
	url := "https://google.com"

	if err := validate_url(url); err == nil {
		t.Errorf("%s should return an error", url)
	}

}

func TestErrorIfNoPath(t *testing.T) {
	url := "https://adventofcode.com"

	if err := validate_url(url); err == nil {
		t.Errorf("%s should return an error", url)
	}
}

func TestErrorIfPathTooShort(t *testing.T) {
	url := "https://adventofcode.com"
	path := "/1/2"

	if err := validate_url(fmt.Sprintf("%s%s", url, path)); err == nil {
		t.Errorf("Should have error because %s is too short", path)
	}
}

func TestErrorIfPathTooLong(t *testing.T) {
	url := "https://adventofcode.com"
	path := "/1/2/3/4"

	if err := validate_url(fmt.Sprintf("%s%s", url, path)); err == nil {
		t.Errorf("Should have error because %s is too long", path)
	}
}

func TestErrorIfNoYear(t *testing.T) {
	url := "http://adventofcode.com/"
	path := "not/a/year"

	if err := validate_url(fmt.Sprintf("%s%s", url, path)); err == nil {
		t.Errorf("Should have error because %s does not have a year", path)
	}
}

func TestErrorOnEarlyYear(t *testing.T) {
	url := "http://adventofcode.com/"
	const year = 2014

	if err := validate_url(fmt.Sprintf("%s%d", url, year)); err == nil {
		t.Errorf("%d should be an invalid year", year)
	}
}

func TestErrorOnLateYear(t *testing.T) {
	url := "http://adventofcode.com/"
	const year = 2023

	if err := validate_url(fmt.Sprintf("%s%d", url, year)); err == nil {
		t.Errorf("%d should be an invalid year", year)
	}
}

func TestErrorIfDayNotInPath(t *testing.T) {
	url := "http://adventofcode.com/2021"
	not_day := "/not/1"

	if err := validate_url(fmt.Sprintf("%s%s", url, not_day)); err == nil {
		t.Error("Should be an error when no day in url", url)
	}
}

func TestErrorIfDayTooLow(t *testing.T) {
	url := "http://adventofcode.com/2021/day/"
	day := 0

	if err := validate_url(fmt.Sprintf("%s%d", url, day)); err == nil {
		t.Errorf("%d should not be a valid day", day)
	}
}

func TestErrorIfDayTooHigh(t *testing.T) {
	url := "http://adventofcode.com/2021/day/"
	day := 26

	if err := validate_url(fmt.Sprintf("%s%d", url, day)); err == nil {
		t.Errorf("%d should not be a valid day", day)
	}
}
