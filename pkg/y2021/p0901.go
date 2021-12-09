package y2021

import (
	"bufio"
	"github.com/timdeklijn/aoc/pkg/solution"
	"strconv"
	"strings"
)

const e91 string = `2199943210
3987894921
9856789892
8767896789
9899965678`

type P0901 struct{}

type Floor map[int]map[int]int

func (p *P0901) GetExamples() solution.Examples {
	return solution.Examples{
		solution.Example{
			N:    1,
			In:   e91,
			Want: 15,
		},
	}
}

func (p *P0901) Run(data *bufio.Scanner) int {

	tot := 0
	row := 0

	// =================================================================
	// Parse input
	// =================================================================

	floor := Floor{}
	for data.Scan() {
		s := data.Text()
		r := map[int]int{}
		for col, num := range strings.Split(s, "") {
			i, _ := strconv.Atoi(num)
			r[col] = i
		}
		floor[row] = r
		row++
	}

	for row := 0; row < len(floor); row++ {
		for col := 0; col < len(floor[0]); col++ {

			// =================================================================
			// Find Neighbours
			// =================================================================

			var nums []int

			// down
			if num, ok := floor[row+1][col]; ok {
				nums = append(nums, num)
			}

			// up
			if num, ok := floor[row-1][col]; ok {
				nums = append(nums, num)
			}

			// left
			if num, ok := floor[row][col-1]; ok {
				nums = append(nums, num)
			}

			// right
			if num, ok := floor[row][col+1]; ok {
				nums = append(nums, num)
			}

			// =================================================================
			// Is it lower?
			// =================================================================

			lower := true
			for _, n := range nums {
				if n <= floor[row][col] {
					lower = !lower
					break
				}
			}
			if lower {
				tot += floor[row][col] + 1
			}
		}
	}
	return tot
}
