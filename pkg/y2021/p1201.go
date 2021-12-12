package y2021

import (
	"bufio"
	"github.com/timdeklijn/aoc/pkg/solution"
	"strings"
	"unicode"
)

const e121 string = `start-A
start-b
A-c
A-b
b-d
A-end
b-end`

const e122 string = `dc-end
HN-start
start-kj
dc-start
dc-HN
LN-dc
HN-end
kj-sa
kj-HN
kj-dc`

type P1201 struct{}

func (p *P1201) GetExamples() solution.Examples {
	return solution.Examples{
		solution.Example{
			N:    1,
			In:   e121,
			Want: 10,
		},
		solution.Example{
			N:    1,
			In:   e122,
			Want: 19,
		},
	}
}

func copyMap(visits map[string]int) map[string]int {
	res := make(map[string]int, len(visits))
	for k, v := range visits {
		res[k] = v
	}
	return res
}

func countRoutes(tst map[string][]string, current string, visited map[string]int) int {
	visited[current] += 1
	sum := 0
	for _, c := range tst[current] {
		if c == "end" {
			sum++
			continue
		}
		if !unicode.IsLower(rune(c[0])) || visited[c] == 0 {
			nextVisited := visited
			if unicode.IsLower(rune(c[0])) {
				nextVisited = copyMap(visited)
			}
			sum += countRoutes(tst, c, nextVisited)
		}
	}
	return sum
}

func (p *P1201) Run(data *bufio.Scanner) int {

	tst := map[string][]string{}
	visited := map[string]int{}

	for data.Scan() {
		s := data.Text()
		spl := strings.Split(s, "-")
		tst[spl[0]] = append(tst[spl[0]], spl[1])
		tst[spl[1]] = append(tst[spl[1]], spl[0])
	}
	return countRoutes(tst, "start", visited)
}
