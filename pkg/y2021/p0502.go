package y2021

import (
	"bufio"

	"github.com/timdeklijn/aoc/pkg/solution"
)

type P0502 struct{}

func (p *P0502) GetExamples() solution.Examples {
	return solution.Examples{
		solution.Example{
			N:    1,
			In:   e51,
			Want: 12,
		},
	}
}

func (p *P0502) Run(data *bufio.Scanner) int {

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
			xdir := Step(dx)
			ydir := Step(dy)
			for i := 0; i < Abs(dx)+1; i++ {
				pipes.AddPipeCoord(x1+(xdir*i), y1+(ydir*i))
			}
		}
	}

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
