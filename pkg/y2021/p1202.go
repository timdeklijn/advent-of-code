package y2021

import (
	"bufio"
	"github.com/timdeklijn/aoc/pkg/solution"
	"strings"
	"unicode"
)

type P1202 struct{}

func (p *P1202) GetExamples() solution.Examples {
	return solution.Examples{
		solution.Example{
			N:    1,
			In:   e121,
			Want: 36,
		},
	}
}

func countRoutesPartTwo(
	tst map[string][]string, current string, visited map[string]int, twice bool,
) int {

	// keep track of how many times a cave has been visited
	if visited[current] > 0 && unicode.IsLower(rune(current[0])) {
		twice = true
	}
	visited[current] += 1

	sum := 0

	for _, c := range tst[current] {

		// Increase routes when encountering an 'end'
		if c == "end" {
			sum++
			continue
		}

		// if there is not a start, not a lower case and the case has not been visited
		// enough: continue
		if c != "start" && (!unicode.IsLower(rune(c[0])) || visited[c] < 1 || !twice) {
			nextVisited := visited
			if unicode.IsLower(rune(c[0])) {
				nextVisited = copyMap(visited)
			}
			sum += countRoutesPartTwo(tst, c, nextVisited, twice)
		}
	}
	return sum
}

func (p *P1202) Run(data *bufio.Scanner) int {

	tst := map[string][]string{}
	visited := map[string]int{}

	for data.Scan() {
		s := data.Text()
		spl := strings.Split(s, "-")
		tst[spl[0]] = append(tst[spl[0]], spl[1])
		tst[spl[1]] = append(tst[spl[1]], spl[0])
	}
	return countRoutesPartTwo(tst, "start", visited, false)
}
