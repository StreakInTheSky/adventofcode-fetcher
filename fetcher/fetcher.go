package fetcher

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

const first_year = 2015
const current_year = 2022

func validate_url(input_url string) error {
	parsed_url, err := url.Parse(input_url)

	if err != nil {
		return err
	}

	if parsed_url.Host != "adventofcode.com" {
		return errors.New(fmt.Sprintf("%s is not a valid advent of code url", input_url))
	}

	parsed_path := strings.Split(parsed_url.Path, "/")

	if len(parsed_path) < 4 {
		return errors.New("Url did not include a day")
	}

	year, err := strconv.Atoi(parsed_path[1])
	if err != nil {
		return err
	}
	if year < first_year || year > current_year {
		return errors.New(fmt.Sprintf("Invalid year: %d", year))
	}

	if parsed_path[2] != "day" {
		return errors.New("Url does not include day")
	}

	day, err := strconv.Atoi(parsed_path[3])
	if err != nil {
		return err
	}
	if day < 1 || day > 25 {
		return errors.New(fmt.Sprintf("%d is not a valid day", day))
	}

	return nil
}
