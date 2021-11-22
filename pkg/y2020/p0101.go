package y2020

import "github.com/timdeklijn/aoc/pkg/solution"

type P0101 struct{}

func (p *P0101) GetExamples() solution.Examples {
	return solution.Examples{solution.Example{
		N:    1,
		In:   "input",
		Want: 3,
	}}
}
func (p *P0101) Run(string) int {
	return 2
}
