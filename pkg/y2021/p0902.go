package y2021

import (
	"bufio"
	"github.com/timdeklijn/aoc/pkg/solution"
	"sort"
	"strconv"
	"strings"
)

type P0902 struct{}

func (p *P0902) GetExamples() solution.Examples {
	return solution.Examples{
		solution.Example{
			N:    1,
			In:   e91,
			Want: 1134,
		},
	}
}

// getNeighbours returns the coordinate of neighbouring (and existing) locations
// on the floor
func getNeighbours(floor Floor, y, x int) [][2]int {

	var nums [][2]int

	// down
	if _, ok := floor[y+1][x]; ok {
		nums = append(nums, [2]int{y + 1, x})
	}

	// up
	if _, ok := floor[y-1][x]; ok {
		nums = append(nums, [2]int{y - 1, x})
	}

	// left
	if _, ok := floor[y][x-1]; ok {
		nums = append(nums, [2]int{y, x - 1})
	}

	// right
	if _, ok := floor[y][x+1]; ok {
		nums = append(nums, [2]int{y, x + 1})
	}

	return nums
}

func (p *P0902) Run(data *bufio.Scanner) int {

	var lowPoints [][2]int
	floor := Floor{}
	row := 0

	// create floor board
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

	// get the lowest points on the floor
	for row := 0; row < len(floor); row++ {
		for col := 0; col < len(floor[0]); col++ {

			crds := getNeighbours(floor, row, col)
			lower := true
			for _, n := range crds {
				if floor[n[0]][n[1]] <= floor[row][col] {
					lower = !lower
					break
				}
			}
			if lower {
				lowPoints = append(lowPoints, [2]int{row, col})
			}
		}
	}

	// check the size of basin for all low points
	var sizes []int
	for _, lp := range lowPoints {
		sizes = append(sizes, basinSize(floor, lp))
	}

	// return the product of the three larges basin sizes
	sort.Ints(sizes)
	return sizes[len(sizes)-3] * sizes[len(sizes)-2] * sizes[len(sizes)-1]
}

// contains checks if a coordinate is in a list of coordinates
func contains(t [2]int, l [][2]int) bool {
	for _, n := range l {
		if n[0] == t[0] && n[1] == t[1] {
			return true
		}
	}
	return false
}

// basinSize returns the size of the basin emanating from the input coordinate.
func basinSize(floor Floor, input [2]int) int {
	crds := [][2]int{input} // collect all coordinates in the basing
	for {
		cnt := 0 // if there are no more neighbours under 9 found we break
		for _, c := range crds {
			n := getNeighbours(floor, c[0], c[1])

			// no neighbours no problems
			if len(n) == 0 {
				continue
			}

			// look for new neighbours lower than 9 and not in crds yet.
			for _, nn := range n {
				if floor[nn[0]][nn[1]] != 9 && !contains(nn, crds) {
					crds = append(crds, nn)
					cnt++
				}
			}
		}

		// break if there are no new neighbours found
		if cnt == 0 {
			break
		}

	}
	return len(crds)
}
