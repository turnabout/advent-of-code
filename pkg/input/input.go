package input

import (
	"embed"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	urlFormat       = "https://adventofcode.com/%d/day/%d/input"
	SessionFileName = "session"
	sessionKeyId    = "session"
)

//go:embed *.txt
var txtFiles embed.FS

// Fetch2021Input fetches problem input for the given day in 2021.
func Fetch2021Input(day int) string {
	// return fetchFileInput(day)
	return strings.TrimSpace(fetchWebInput(2021, day))
}

func fetchFileInput(day int) string {
	b, err := txtFiles.ReadFile(fmt.Sprintf("%d.txt", day))
	if err != nil {
		log.Fatal(err)
	}

	return string(b)
}

// fetchWebInput fetches problem input for the given year and day, from the
// original web page.
func fetchWebInput(year int, day int) string {

	// Create request
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf(urlFormat, year, day),
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Include session header in request
	req.Header = http.Header{
		"cookie": {
			fmt.Sprintf("%s=%s", sessionKeyId, getSession()),
		},
	}

	// Launch request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	// Return response as a string
	content, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}

func getSession() string {
	if _, err := os.Stat(SessionFileName); os.IsNotExist(err) {
		log.Fatalf("File containing session data (%s) does not exist", SessionFileName)
	}

	content, err := os.ReadFile(SessionFileName)
	if err != nil {
		log.Fatal(err)
	}

	return strings.TrimSpace(string(content))
}
