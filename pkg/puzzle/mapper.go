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
	"20210201": &y2021.P0201{},
	"20210202": &y2021.P0202{},
	"20210301": &y2021.P0301{},
	"20210302": &y2021.P0302{},
	"20210401": &y2021.P0401{},
	"20210402": &y2021.P0402{},
	"20210501": &y2021.P0501{},
	"20210502": &y2021.P0502{},
	"20210601": &y2021.P0601{},
	"20210602": &y2021.P0602{},
	"20210701": &y2021.P0701{},
	"20210702": &y2021.P0702{},
	"20210801": &y2021.P0801{},
	"20210802": &y2021.P0802{},
	"20210901": &y2021.P0901{},
	"20210902": &y2021.P0902{},
	"20211001": &y2021.P1001{},
	"20211002": &y2021.P1002{},
	"20211101": &y2021.P1101{},
	"20211102": &y2021.P1102{},
	"20211201": &y2021.P1201{},
	"20211202": &y2021.P1202{},
	"20211301": &y2021.P1301{},
	"20211302": &y2021.P1302{},
	"20211401": &y2021.P1401{},
	"20211402": &y2021.P1402{},
	"20211501": &y2021.P1501{},
	"20211502": &y2021.P1502{},
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
