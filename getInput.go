package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// getInput gets the input for a given year and day's puzzle.
func getInput(year int, day int) string {
	// Parse input URL
	u, err := url.Parse(
		fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day),
	)
	if err != nil {
		log.Fatalf("failed to parse URL for input: %w", err)
	}

	// Create request
	req := http.Request{
		Method: "GET",
		URL:    u,
		Header: http.Header{
			"Cookie": []string{
				"session=" + os.Getenv("SESSION"),
			},
		},
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
	}

	// Send
	res, err := http.DefaultClient.Do(&req)
	if err != nil {
		log.Fatalf("failed to get response for input: %w", err)
	}

	// Read contents
	bs, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("failed to get read response contents for input: %w", err)
	}

	return strings.TrimSpace(string(bs))
}
