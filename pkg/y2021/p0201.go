package y2021

import (
	"bufio"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/timdeklijn/aoc/pkg/solution"
)

const e21 string = `forward 5
down 5
forward 8
up 3
down 8
forward 2`

type P0201 struct{}

func (p *P0201) GetExamples() solution.Examples {
	return solution.Examples{
		solution.Example{
			N:    1,
			In:   e21,
			Want: 150,
		},
	}
}

func (p *P0201) Run(data *bufio.Scanner) int {

	// read scanner to list of ints
	depth := 0
	pos := 0
	for data.Scan() {
		s := data.Text()

		spl := strings.Split(s, " ")

		dist, err := strconv.Atoi(spl[1])
		if err != nil {
			log.Panicf("error converting '%s' to int: %e", spl[1], err)
		}

		switch spl[0] {
		case "forward":
			pos += dist
		case "down":
			depth += dist
		case "up":
			depth -= dist
		}
	}

	return pos * depth
}
