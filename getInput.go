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
// It will first attempt to recover it from cache. If no cache exists, it delegates to getSourceInput().
func getInput(year int, day int) string {
	// Recover cached input, if it exists
	cachedInput := getCachedInput(year, day)
	if cachedInput != "" {
		return cachedInput
	}

	// Get input from source
	input := getSourceInput(year, day)

	// Set in cache
	setInputCache(input, year, day)

	return input
}

// getSourceInput gets the input for a given year and day's puzzle, from the source website.
func getSourceInput(year int, day int) string {
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

// getCachedFilePath gets the cache file name for the given year and day.
func getCachedFilePath(year int, day int) string {
	return fmt.Sprintf("cache/%d_%d.txt", year, day)
}

// getCachedInput recovers the cached input for the given year and day.
func getCachedInput(year int, day int) string {
	filePath := getCachedFilePath(year, day)
	if !fileExists(filePath) {
		return ""
	}

	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("failed to open cached input file: %w", err)
	}
	defer f.Close()

	bs, err := io.ReadAll(f)
	if err != nil {
		log.Fatalf("failed to read cached input file: %w", err)
	}

	return string(bs)
}

// setInputCache sets the input cache for the given year and day.
func setInputCache(input string, year int, day int) {
	file, err := os.Create(getCachedFilePath(year, day))
	if err != nil {
		log.Fatalf("failed to create cached input file")
	}
	defer file.Close()

	if _, err := file.WriteString(input); err != nil {
		log.Fatalf("failed to write input into cache file: %w")
	}
}

// fileExists checks if the given file exists.
func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}
