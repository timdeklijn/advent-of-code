package y2021

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/timdeklijn/aoc/pkg/solution"
)

type P0701 struct{}

const e71 string = "16,1,2,0,4,2,7,1,2,14"

func (p *P0701) GetExamples() solution.Examples {
	return solution.Examples{
		solution.Example{
			N:    1,
			In:   e71,
			Want: 37,
		},
	}
}

// simpleCrabAlign returns the position that is cheapest to align to
func simpleCrabAlign(input []int) int {
	fuel := 1000000
	for pos := range input {
		sum := 0
		for _, el := range input {
			sum += Abs(el - pos)
		}
		if sum < fuel {
			fuel = sum
		}

	}
	return fuel
}

func (p *P0701) Run(data *bufio.Scanner) int {

	var input []int
	for data.Scan() {
		s := data.Text()
		spl := strings.Split(s, ",")
		for _, i := range spl {
			ii, _ := strconv.Atoi(i)
			input = append(input, ii)
		}
	}

	return simpleCrabAlign(input)
}
