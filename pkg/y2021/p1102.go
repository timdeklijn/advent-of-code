package y2021

import (
	"bufio"
	"github.com/timdeklijn/aoc/pkg/solution"
	"strconv"
	"strings"
)

type P1102 struct{}

func (p *P1102) GetExamples() solution.Examples {
	return solution.Examples{
		solution.Example{
			N:    1,
			In:   e111,
			Want: 195,
		},
	}
}

// fullFlash checks if the whole 'Octo' is 0
func (oc Octos) fullFlash() bool {
	for r, _ := range oc {
		for c, _ := range oc[0] {
			if oc[r][c] != 0 {
				return false
			}
		}
	}
	return true
}

func (p *P1102) Run(data *bufio.Scanner) int {

	// create a 2D map of octopuses from the input
	octos := Octos{}
	row := 0
	for data.Scan() {
		s := data.Text()
		spl := strings.Split(s, "")
		tmp := map[int]int{}
		for col, n := range spl {
			i, _ := strconv.Atoi(n)
			tmp[col] = i
		}
		octos[row] = tmp
		row++
	}

	// find the step that all octo's have flashed at the same time
	step := 0
	for {
		octos.step()
		step++
		if octos.fullFlash() {
			break
		}
	}

	return step
}
