package y2020

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/timdeklijn/aoc/pkg/solution"
)

const e1 string = `1721
979
366
299
675
1456`

type P0101 struct{}

func (p *P0101) GetExamples() solution.Examples {
	return solution.Examples{
		solution.Example{
			N:    1,
			In:   e1,
			Want: 514579,
		},
	}
}

func (p *P0101) Run(data *bufio.Scanner) int {

	l := []int{}
	for data.Scan() {
		s := data.Text()

		// TODO: this should be a utils function
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Panicf("error converting '%s' to int: %e", s, err)
		}

		l = append(l, i)
	}
	if err := data.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	for _, i := range l {
		for _, j := range l {
			if i+j == 2020 {
				return i * j
			}
		}
	}
	return 0
}
