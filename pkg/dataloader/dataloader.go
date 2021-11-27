package dataloader

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

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

const dataPath string = "/Users/timdeklijn/go/src/github.com/timdeklijn/aoc/data/"

func createPath(day, year int) (string, error) {
	if _, err := os.Stat(dataPath); os.IsNotExist(err) {
		return "", fmt.Errorf("datapath '%s' does not exist", dataPath)
	}
	yearPath := fmt.Sprintf("%s%d/", dataPath, year)
	if _, err := os.Stat(yearPath); os.IsNotExist(err) {
		log.Infof("creating datafolder: %s", yearPath)
		if err := os.Mkdir(yearPath, 0755); err != nil {
			return "", fmt.Errorf("error creating datafolder '%s': %e", yearPath, err)
		}
	}

	var d string
	if day < 9 {
		d = fmt.Sprintf("%02d", day)
	} else {
		d = strconv.Itoa(day)
	}
	dayPath := fmt.Sprintf("%s/%s.txt", yearPath, d)

	return dayPath, nil
}

func NewDataLoader(day, year int) (DataLoader, error) {
	path, err := createPath(day, year)
	log.Infof("datapath: %s", path)
	if err != nil {
		return DataLoader{}, err
	}
	return DataLoader{path, Puzzle{day, year}}, nil
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

// getData sends an http GET request to the advent of code server to retrieve
// data. A cookie (retrieved as env variable) is added to the header of the GET
// request.
func (d *DataLoader) getData() ([]byte, error) {
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

	return body, nil
}

func readerFromFile(f string) (*bufio.Scanner, error) {
	file, err := os.Open(f)
	if err != nil {
		return nil, err
	}

	return bufio.NewScanner(bufio.NewReader(file)), nil
}

// RetrieveData loads the data from a specific url. To succesfully do a GET
// requests a cookie is constructed from a session id.
func (d *DataLoader) RetrieveData() (*bufio.Scanner, error) {
	if _, err := os.Stat(d.Path); err == nil {
		log.Info("data already exists")
		reader, err := readerFromFile(d.Path)
		if err != nil {
			return nil, err
		}
		return reader, nil
	}

	data, err := d.getData()
	if err != nil {
		return nil, err
	}

	f, err := os.Create(d.Path)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	if _, err := f.Write(data); err != nil {
		return nil, err
	}

	reader, err := readerFromFile(d.Path)
	if err != nil {
		return nil, err
	}
	return reader, nil
}
