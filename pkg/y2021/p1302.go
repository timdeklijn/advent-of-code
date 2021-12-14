package y2021

import (
	"bufio"
	"github.com/timdeklijn/aoc/pkg/solution"
	"strconv"
	"strings"
)

type P1302 struct{}

func (p *P1302) GetExamples() solution.Examples {
	return solution.Examples{
		solution.Example{
			N:    1,
			In:   e131,
			Want: 16,
		},
	}
}

func (p *P1302) Run(data *bufio.Scanner) int {

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

	for _, fold := range folds {
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

	instructions.print()

	return sum
}
