package puzzle

import (
	"fmt"
	"strconv"

	"github.com/timdeklijn/aoc/pkg/solution"
	"github.com/timdeklijn/aoc/pkg/y2020"
)

var solutionMap = map[string]solution.Solution{
	"20200101": &y2020.P0101{},
}

func toKey(y, d, p int) string {
	yS := strconv.Itoa(y)
	var dS string
	if d < 10 {
		dS = fmt.Sprintf("%02d", d)
	}
	dS = strconv.Itoa(d)
	pS := fmt.Sprintf("%02d", p)
	return yS + dS + pS
}

func Map(y, d, p int) solution.Solution {
	// TODO: add return when solution is not (yet) implemented
	key := toKey(y, d, p)
	return solutionMap[key]
}
