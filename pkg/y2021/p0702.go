package y2021

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/timdeklijn/aoc/pkg/solution"
)

type P0702 struct{}

func (p *P0702) GetExamples() solution.Examples {
	return solution.Examples{
		solution.Example{
			N:    1,
			In:   e71,
			Want: 168,
		},
	}
}

//advancedCrabAlign calculates the fuel consumption for advanced alignment
func advancedCrabAlign(input []int) int {
	fuel := 10000000000
	for pos := range input {
		sum := 0
		for _, el := range input {
			for i := 1; i < Abs(el-pos)+1; i++ {
				sum += i
			}
		}
		if sum < fuel {
			fuel = sum
		}
	}
	return fuel
}

func (p *P0702) Run(data *bufio.Scanner) int {

	var input []int
	for data.Scan() {
		s := data.Text()
		spl := strings.Split(s, ",")
		for _, i := range spl {
			ii, _ := strconv.Atoi(i)
			input = append(input, ii)
		}
	}

	return advancedCrabAlign(input)
}
