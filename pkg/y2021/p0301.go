package y2021

import (
	"bufio"
	"strconv"

	"github.com/timdeklijn/aoc/pkg/solution"
)

const e31 string = `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`

type P0301 struct{}

func (p *P0301) GetExamples() solution.Examples {
	return solution.Examples{
		solution.Example{
			N:    1,
			In:   e31,
			Want: 198,
		},
	}
}

func (p *P0301) Run(data *bufio.Scanner) int {

	// read scanner to list of ints
	lines := []string{}
	for data.Scan() {
		s := data.Text()
		lines = append(lines, s)
	}

	gamma := ""
	epsilon := ""
	for i := 0; i < len(lines[0]); i++ {
		ones := 0
		zeros := 0
		for _, l := range lines {
			switch l[i] {
			case '0':
				zeros += 1
			case '1':
				ones += 1
			}
		}

		if zeros > ones {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}
	}

	g, _ := strconv.ParseInt(gamma, 2, 64)
	e, _ := strconv.ParseInt(epsilon, 2, 64)

	return int(g * e)
}
