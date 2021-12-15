package y2021

import (
	"bufio"
	"container/heap"
	"strconv"
	"strings"

	"github.com/timdeklijn/aoc/pkg/solution"
)

const e151 string = `1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581`

type P1501 struct{}

func (p *P1501) GetExamples() solution.Examples {
	return solution.Examples{
		solution.Example{
			N:    1,
			In:   e151,
			Want: 40,
		},
	}
}

type qi struct {
	pos       xy
	riskLevel int
	index     int
}

type PriorityQueue []qi

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].riskLevel < pq[j].riskLevel
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(qi)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

var dx = [4]int{0, 0, -1, 1}

var dy = [4]int{-1, 1, 0, 0}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

type xy struct{ x, y int }

type Ceiling map[xy]int

func (p *P1501) Run(data *bufio.Scanner) int {

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
	target := xy{maxx, maxy}

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
			risk = risk - 1
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

			if next.x >= maxx+1 || next.x < 0 || next.y > maxy+1 || next.y < 0 {
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
