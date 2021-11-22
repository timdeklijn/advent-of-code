package puzzle

import (
	"github.com/timdeklijn/aoc/pkg/solution"
	"github.com/timdeklijn/aoc/pkg/y2020"
)

var solutionMap = map[string]solution.Solution{
	"20200101": &y2020.P0101{},
}

// TODO: add function to convert y,d,p to YYYYDDPP string
func Map(y, d, p int) solution.Solution {
	// TODO: add return when solution is not (yet) implemented
	return solutionMap["20200101"]
}
