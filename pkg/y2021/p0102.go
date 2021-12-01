package y2021

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/timdeklijn/aoc/pkg/solution"
)

const e12 string = `199
200
208
210
200
207
240
269
260
263`

type P0102 struct{}

func (p *P0102) GetExamples() solution.Examples {
	return solution.Examples{
		solution.Example{
			N:    1,
			In:   e12,
			Want: 5,
		},
	}
}

func (p *P0102) Run(data *bufio.Scanner) int {

	// read scanner to list of ints
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

	// create list of sums over a window of three
	windows := []int{}
	for i := 0; i < len(l)-2; i++ {
		windows = append(windows, l[i]+l[i+1]+l[i+2])
	}

	// count the number of times a number is larger than the previous number.
	c := 0
	for i := 1; i < len(windows); i++ {
		if windows[i] > windows[i-1] {
			c++
		}

	}
	return c
}
