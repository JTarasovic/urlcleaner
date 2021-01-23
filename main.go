package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Head(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	u := resp.Request.URL
	fmt.Fprintf(os.Stdout, "%s://%s%s", u.Scheme, u.Host, u.Path)
}
