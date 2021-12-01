package puzzle

import (
	"fmt"
	"strconv"

	"github.com/timdeklijn/aoc/pkg/solution"
	"github.com/timdeklijn/aoc/pkg/y2020"
	"github.com/timdeklijn/aoc/pkg/y2021"
)

var solutionMap = map[string]solution.Solution{
	"20200101": &y2020.P0101{},
	"20210101": &y2021.P0101{},
	"20210102": &y2021.P0102{},
}

func toKey(y, d, p int) string {
	yS := strconv.Itoa(y)
	var dS string
	if d < 10 {
		dS = fmt.Sprintf("%02d", d)
	} else {
		dS = strconv.Itoa(d)
	}
	pS := fmt.Sprintf("%02d", p)
	return yS + dS + pS
}

func Map(y, d, p int) (solution.Solution, error) {
	key := toKey(y, d, p)
	if sol, ok := solutionMap[key]; ok {
		return sol, nil
	}
	return nil, fmt.Errorf("could not find solution: %s", key)
}
