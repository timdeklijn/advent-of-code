package y2021

import (
	"bufio"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/timdeklijn/aoc/pkg/solution"
	"strconv"
	"strings"
)

const e111 string = `5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`

type P1101 struct{}

func (p *P1101) GetExamples() solution.Examples {
	return solution.Examples{
		solution.Example{
			N:    1,
			In:   e111,
			Want: 1656,
		},
	}
}

// Octos is a 2D map with octopuses that have an energy level
type Octos map[int]map[int]int

// initOctos initiates Octo with a specific value in all fields
func initOctos(r, c, n int) Octos {
	oc := Octos{}
	for rr := 0; rr < r; rr++ {
		tmp := map[int]int{}
		for cc := 0; cc < c; cc++ {
			tmp[cc] = n
		}
		oc[rr] = tmp
	}
	return oc
}

// print pretty prints all octopuses
func (oc Octos) print() {
	for r := 0; r < len(oc); r++ {
		s := ""
		for c := 0; c < len(oc[0]); c++ {
			s += fmt.Sprintf("%d", oc[r][c])
		}
		logrus.Infof("\t%s", s)
	}
}

// increase adds 1 to all neighbours of r,c
func increase(oc Octos, r, c int) Octos {
	for y := -1; y < 2; y++ {
		for x := -1; x < 2; x++ {
			if _, ok := oc[r+y][c+x]; ok {
				oc[r+y][c+x]++
			}
		}
	}
	return oc
}

// add simply sums the two input octo's
func add(oc, add Octos) Octos {
	for r, _ := range oc {
		for c, _ := range oc[0] {
			oc[r][c] += add[r][c]
		}
	}
	return oc
}

// step adds 1 to all spots and then starts checking for values over 9. If found,
// increase its neighbours by a second '1'. Return the number of times a value
// over 9 has been found.
func (oc Octos) step() int {
	oc = add(oc, initOctos(len(oc), len(oc[0]), 1))
	var flashed [][2]int

	for {
		newOcto := initOctos(len(oc), len(oc[0]), 0)
		changed := false

		for r, _ := range oc {
			for c, _ := range oc[0] {
				if oc[r][c] > 9 {
					newOcto = increase(newOcto, r, c)
					flashed = append(flashed, [2]int{r, c})
					changed = true
				}
			}
		}

		// everything that flashed should be set to 0
		oc = add(oc, newOcto)
		for _, c := range flashed {
			oc[c[0]][c[1]] = 0
		}

		// If there are no changes we finished for this board
		if !changed {
			break
		}
	}

	return len(flashed)
}

func (p *P1101) Run(data *bufio.Scanner) int {

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

	// count the number of flashes over a 100 steps.
	sum := 0
	for i := 0; i < 100; i++ {
		n := octos.step()
		sum += n
	}

	return sum
}
