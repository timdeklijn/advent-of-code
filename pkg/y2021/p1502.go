package y2021

import (
	"bufio"
	"container/heap"
	"strconv"
	"strings"

	"github.com/timdeklijn/aoc/pkg/solution"
)

type P1502 struct{}

func (p *P1502) GetExamples() solution.Examples {
	return solution.Examples{
		solution.Example{
			N:    1,
			In:   e151,
			Want: 315,
		},
	}
}

func (p *P1502) Run(data *bufio.Scanner) int {

	var maxx, maxy = 0, 0
	ceiling := Ceiling{}
	y := 0

	for data.Scan() {
		s := data.Text()
		spl := strings.Split(s, "")
		for x, i := range spl {
			ii, _ := strconv.Atoi(i)
			ceiling[xy{x, y}] = ii
			maxx = max(maxx, x)
			maxy = max(maxy, y)
		}
		y++
	}

	start := xy{0, 0}
	target := xy{(maxx+1)*5 - 1, (maxy+1)*5 - 1}

	risk := func(pos xy) int {
		og := xy{
			pos.x % (maxx + 1),
			pos.y % (maxy + 1),
		}

		mx := (pos.x) / (maxx + 1)
		my := (pos.y) / (maxy + 1)

		mul := mx + my

		risk := ceiling[og] + mul
		if risk > 9 {
			risk = risk - 9
		}
		return risk
	}

	shortestAt := map[xy]int{}

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	pq.Push(qi{pos: start, riskLevel: 0})

	for pq.Len() > 0 {
		head := heap.Pop(&pq).(qi)

		for i := 0; i < 4; i++ {
			next := xy{
				head.pos.x + dx[i],
				head.pos.y + dy[i],
			}

			if next.x >= (maxx+1)*5 || next.x < 0 || next.y > (maxy+1)*5 || next.y < 0 {
				continue
			}

			nextRisk := head.riskLevel + risk(next)

			if sAt, ok := shortestAt[next]; ok && sAt <= nextRisk {
				continue
			} else {
				shortestAt[next] = nextRisk
			}

			pq.Push(qi{
				pos:       next,
				riskLevel: nextRisk,
			})

		}
	}

	return shortestAt[target]
}
