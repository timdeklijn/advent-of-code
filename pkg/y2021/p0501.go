package y2021

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/timdeklijn/aoc/pkg/solution"
)

const e51 string = `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`

type P0501 struct{}

func (p *P0501) GetExamples() solution.Examples {
	return solution.Examples{
		solution.Example{
			N:    1,
			In:   e51,
			Want: 5,
		},
	}
}

// SplitPipeLine returns the xy of the start and end coordinate
func SplitPipeLine(s string) (int, int, int, int) {
	spl := strings.Split(s, " -> ")

	spl1 := strings.Split(spl[0], ",")
	spl2 := strings.Split(spl[1], ",")

	x1, _ := strconv.Atoi(spl1[0])
	x2, _ := strconv.Atoi(spl2[0])
	y1, _ := strconv.Atoi(spl1[1])
	y2, _ := strconv.Atoi(spl2[1])
	return x1, y1, x2, y2
}

// Pipes holds the area that pipes are laid on and can be queried using p[y][x]
type Pipes map[int]map[int]int

// AddPipeCoord will add 1 to a coordinate in the pipe area
func (p Pipes) AddPipeCoord(x, y int) {
	if _, ok := p[y]; !ok {
		p[y] = map[int]int{}
		p[y][x] = 1
	} else {
		p[y][x] += 1
	}
}

// Abs returns the absolute value of the input integer
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Step returns -1 or 1 depending on the sign of the input
func Step(x int) int {
	if x < 0 {
		return -1
	}
	return 1
}

func (p *P0501) Run(data *bufio.Scanner) int {

	pipes := make(Pipes)

	for data.Scan() {
		s := data.Text()
		x1, y1, x2, y2 := SplitPipeLine(s)

		dx := x2 - x1
		dy := y2 - y1

		// vertical line
		if dx == 0 {
			dir := Step(dy)
			for i := 0; i < Abs(dy)+1; i++ {
				pipes.AddPipeCoord(x1, y1+(dir*i))
			}
			// horizontal
		} else if dy == 0 {
			dir := Step(dx)
			for i := 0; i < Abs(dx)+1; i++ {
				pipes.AddPipeCoord(x1+(dir*i), y1)
			}
			// diagonal
		} else {
			continue
		}
	}

	// Count where pipes has a value of two or higher
	sum := 0
	for y, el := range pipes {
		for x := range el {
			if pipes[y][x] > 1 {
				sum++
			}
		}
	}

	return sum
}
