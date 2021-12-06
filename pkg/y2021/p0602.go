package y2021

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/timdeklijn/aoc/pkg/solution"
)

type P0602 struct{}

func (p *P0602) GetExamples() solution.Examples {
	return solution.Examples{
		solution.Example{
			N:    1,
			In:   e61,
			Want: 26984457539,
		},
	}
}

type FishPopulation map[int]int

// update will add the number of 0's to 8 and 6 and reduce the number of the
// other fish
func (p FishPopulation) update() FishPopulation {
	tmp := make(FishPopulation)
	tmp[8] = p[0]
	tmp[6] += p[0]
	for i := 0; i < 8; i++ {
		tmp[i] += p[i+1]
	}
	return tmp
}

func (p *P0602) Run(data *bufio.Scanner) int {

	population := make(FishPopulation)
	days := 256

	for data.Scan() {
		s := data.Text()
		spl := strings.Split(s, ",")
		for _, e := range spl {
			i, _ := strconv.Atoi(e)
			population[i] += 1
		}
	}

	// update population days times
	for i := 0; i < days; i++ {
		population = population.update()
	}

	// Sum all values in the population map
	sum := 0
	for _, v := range population {
		sum += v
	}

	return sum
}
