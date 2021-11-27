package dataloader

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

type Puzzle struct {
	Day  int
	Year int
}

type DataLoader struct {
	Path   string
	Puzzle Puzzle
}

func createPath(day, year int) string {
	// TODO: create a sane path from some base path and the year/day
	// path/to/this/repo/pkg/year<year>/day<day>/
	return "path"
}

func NewDataLoader(day, year int) DataLoader {
	path := createPath(day, year)
	return DataLoader{path, Puzzle{day, year}}
}

// CreateURL constructs an URL for the advent of code dataset for a specific
// year/day combination.
func (d *DataLoader) CreateURL() string {
	return fmt.Sprintf(
		"https://adventofcode.com/%d/day/%d/input",
		d.Puzzle.Year,
		d.Puzzle.Day,
	)
}

// FetchData loads the data from a specific url. To succesfully do a GET
// requests a cookie is constructed from a session id.
func (d *DataLoader) FetchData() ([]byte, error) {
	cookie := os.Getenv("COOKIE_SESSION")
	if cookie == "" {
		return []byte{}, fmt.Errorf("error retrieving cookie from environment")
	}
	url := d.CreateURL()
	log.Infof("fetching data from: %s", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []byte{}, err
	}

	// Create a cookie from the session ID
	req.Header.Set("Cookie", fmt.Sprintf("session=%s", cookie))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return []byte{}, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return body, err
}
