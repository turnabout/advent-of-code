package input

import (
	"embed"
	"fmt"
	"io"
	"log"
	"net/http"
)

const urlFormat = "https://adventofcode.com/%d/day/%d/input"

//go:embed *.txt
var txtFiles embed.FS

// Fetch2021Input fetches problem input for the given day in 2021.
func Fetch2021Input(day int) string {
	return fetchFileInput(day)
	// return fetchWebInput(2021, day)
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
// TODO: User authentication will be needed
func fetchWebInput(year int, day int) string {
	res, err := http.Get(fmt.Sprintf(urlFormat, year, day))
	if err != nil {
		log.Fatal(err)
	}

	content, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}
