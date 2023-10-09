package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const httpPrefix = "http://"

func fetch(url string) {
	addHttpPrefixIfNotProvided(&url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}

	_, err = io.Copy(os.Stdout, resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}
}

func addHttpPrefixIfNotProvided(url *string) {
	if !strings.HasPrefix(*url, httpPrefix) {
		*url = httpPrefix + *url
	}
}
