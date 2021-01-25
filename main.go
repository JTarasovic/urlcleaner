package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
)

func main() {
	// TODO: better argument handling / flag parsing
	out, err := cleanURL(os.Args[1])
	if err != nil {
		// TODO: better error handling
		panic(err)
	}

	fmt.Fprint(os.Stdout, out)
	os.Exit(0)
}

func cleanURL(s string) (string, error) {
	parsed, err := url.Parse(s)
	if err != nil {
		return "", err
	}
	resp, err := http.Head(formatURL(parsed))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	return formatURL(resp.Request.URL), nil
}

func formatURL(u *url.URL) string {
	return fmt.Sprintf("%s://%s%s", u.Scheme, u.Host, u.Path)
}
