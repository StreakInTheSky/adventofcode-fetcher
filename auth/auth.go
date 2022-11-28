package auth

import (
	"context"
	"errors"
	"fmt"

	"github.com/chromedp/chromedp"
)

const (
	Github Service = iota
	Twitter
	Google
)

func (s Service) toString() {
	switch s {
	case Github:
		return "github"
	case Twitter:
		return "twitter"
	case Google:
		return "google"
	}
}





const base_url = "http://adventofcode.com/auth/"

func connect(url string) (context.Context, error) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	res, err := chromedp.RunResponse(ctx, chromedp.Navigate(url))
	if err != nil {
		panic(err)
	}
	if res.Status >= 400 {
		return ctx, errors.New(fmt.Sprintf("%d: %s", res.Status, res.StatusText))
	}

	println(fmt.Printf("%s", res.URL))
	return ctx, nil
}

func type_input(ctx context.Context, query_string, input_string string) error {
	return nil
}

func main() {
	service = Github
	ctx, err := connect("https://adventofcode.com/auth/github")
	if err != nil {
		panic(err)
	}

	ctx, err := type
}
