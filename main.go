package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/url"
	"os"

	"github.com/gobwas/glob"
)

const (
	defaultGlob = "*"
	defaultURL  = "-"
)

var globFlag, urlFlag string

func main() {
	flag.StringVar(&globFlag, "glob", defaultGlob, "the glob used to determine which query params to filter/remove")
	flag.StringVar(&urlFlag, "url", defaultURL, "the URL to filter. `-` for stdin")
	flag.Parse()

	urlToUse := urlFlag
	if urlFlag == defaultURL {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			panic(err)
		}
		urlToUse = scanner.Text()
	}

	u, g, err := prepare(urlToUse, globFlag)
	if err != nil {
		panic(err)
	}
	filterQueryParams(u, g)
	fmt.Fprintln(os.Stdout, u)
}

func prepare(s, glb string) (*url.URL, glob.Glob, error) {
	u, err := url.Parse(s)
	if err != nil {
		return nil, nil, err
	}

	g, err := glob.Compile(globFlag)
	if err != nil {
		return nil, nil, err
	}
	return u, g, nil
}

func filterQueryParams(u *url.URL, g glob.Glob) {
	q := u.Query()
	for k := range q {
		if g.Match(k) {
			q.Del(k)
		}
	}
	// set the raw query to the filtered set of query params
	u.RawQuery = q.Encode()
}
