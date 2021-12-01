package y2021

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/timdeklijn/aoc/pkg/solution"
)

const e1 string = `199
200
208
210
200
207
240
269
260
263`

type P0101 struct{}

func (p *P0101) GetExamples() solution.Examples {
	return solution.Examples{
		solution.Example{
			N:    1,
			In:   e1,
			Want: 7,
		},
	}
}

func (p *P0101) Run(data *bufio.Scanner) int {

	l := []int{}
	for data.Scan() {
		s := data.Text()

		i, err := strconv.Atoi(s)
		if err != nil {
			log.Panicf("error converting '%s' to int: %e", s, err)
		}

		l = append(l, i)
	}

	if err := data.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	c := 0
	for i := 1; i < len(l); i++ {
		if l[i] > l[i-1] {
			c++
		}

	}
	return c
}
