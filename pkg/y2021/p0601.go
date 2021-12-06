package y2021

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/timdeklijn/aoc/pkg/solution"
)

const e61 string = "3,4,3,1,2"

type P0601 struct{}

func (p *P0601) GetExamples() solution.Examples {
	return solution.Examples{
		solution.Example{
			N:    1,
			In:   e61,
			Want: 5934,
		},
	}
}

func (p *P0601) Run(data *bufio.Scanner) int {

	population := []int{}
	days := 80

	for data.Scan() {
		s := data.Text()
		spl := strings.Split(s, ",")
		for _, e := range spl {
			i, _ := strconv.Atoi(e)
			population = append(population, i)
		}
	}

	for i := 0; i < days; i++ {
		n := 0
		for i := range population {
			if population[i] == 0 {
				n++
				population[i] = 6
			} else {
				population[i] = population[i] - 1
			}
		}

		for j := 0; j < n; j++ {
			population = append(population, 8)
		}
	}

	return len(population)
}
