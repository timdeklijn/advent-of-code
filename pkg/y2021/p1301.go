package y2021

import (
	"bufio"
	"github.com/sirupsen/logrus"
	"github.com/timdeklijn/aoc/pkg/solution"
	"strconv"
	"strings"
)

type P1301 struct{}

const e131 string = `6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5`

func (p *P1301) GetExamples() solution.Examples {
	return solution.Examples{
		solution.Example{
			N:    1,
			In:   e131,
			Want: 17,
		},
	}
}

type Instructions map[int]map[int]bool

func (i Instructions) print() {
	maxX := 0
	maxY := 0
	for y, row := range i {
		for x := range row {
			if _, ok := i[y][x]; ok {
				if x > maxX {
					maxX = x
				}
				if y > maxY {
					maxY = y
				}
			}
		}
	}

	logrus.Info()
	for y := 0; y < maxY+1; y++ {
		s := ""
		for x := 0; x < maxX+1; x++ {
			if _, ok := i[y][x]; ok {
				s += "#"
			} else {
				s += "."
			}
		}
		logrus.Info(s)
	}
}

type Fold struct {
	dir string
	pos int
}

func (p *P1301) Run(data *bufio.Scanner) int {

	// =================================================================================
	// PARSE
	// =================================================================================

	instructions := Instructions{}
	var folds []Fold
	isFolds := false
	for data.Scan() {
		s := data.Text()

		if s == "" {
			isFolds = true
			continue
		}

		if !isFolds {
			spl := strings.Split(s, ",")
			x, _ := strconv.Atoi(spl[0])
			y, _ := strconv.Atoi(spl[1])
			if _, ok := instructions[y]; !ok {
				tmp := map[int]bool{}
				tmp[x] = true
				instructions[y] = tmp
			} else {
				instructions[y][x] = true
			}
		} else {
			spl := strings.Split(s[11:], "=")
			pos, _ := strconv.Atoi(spl[1])
			folds = append(folds, Fold{spl[0], pos})
		}
	}

	// =================================================================================
	// FOLD
	// =================================================================================

	for _, fold := range folds[:1] {
		if fold.dir == "x" {
			for y, row := range instructions {
				for x := range row {
					if _, ok := instructions[y][x]; ok {
						if x > fold.pos {
							newX := fold.pos - (x - fold.pos)
							instructions[y][newX] = true
							delete(instructions[y], x)
						}
					}
				}
			}
		} else {
			for y, row := range instructions {
				for x := range row {
					if _, ok := instructions[y][x]; ok {
						if y > fold.pos {
							newY := fold.pos - (y - fold.pos)
							if _, ok := instructions[newY]; ok {
								instructions[newY][x] = true
							} else {
								tmp := map[int]bool{}
								tmp[x] = true
								instructions[newY] = tmp
							}
							delete(instructions[y], x)
						}
					}
				}
			}
		}
	}

	sum := 0
	for y := range instructions {
		sum += len(instructions[y])
	}

	return sum
}
