package y2021

import (
	"bufio"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/timdeklijn/aoc/pkg/solution"
)

type P1402 struct{}

func (p *P1402) GetExamples() solution.Examples {
	return solution.Examples{
		solution.Example{
			N:    1,
			In:   e141,
			Want: 2188189693529,
		},
	}
}

// applyManyRules updates the pairs in count based on the input rules
func applyManyRules(count map[string]int, rules map[string]string) map[string]int {
	newCount := map[string]int{}
	for k, v := range count {
		extra := rules[k]
		p1 := string(k[0]) + extra
		p2 := extra + string(k[1])
		newCount[p1] += v
		newCount[p2] += v
	}
	return newCount
}

func (p *P1402) Run(data *bufio.Scanner) int {

	// ===========================================================================
	// PARSE
	// ===========================================================================

	var templateString string
	rules := map[string]string{}
	temp := true
	for data.Scan() {
		s := data.Text()
		if s == "" {
			temp = false
			continue
		}

		if temp {
			templateString = s
			continue
		}

		spl := strings.Split(s, " -> ")
		rules[spl[0]] = spl[1]
	}

	template := map[string]int{}
	for i := 0; i < len(templateString)-1; i++ {
		pair := templateString[i : i+2]
		logrus.Info(pair)
		template[pair] += 1
	}

	// ===========================================================================
	// APPLY THE RULES
	// ===========================================================================

	for i := 0; i < 40; i++ {
		template = applyManyRules(template, rules)
	}

	logrus.Info(template)
	counts := map[string]int{}
	for k, v := range template {
		counts[string(k[0])] += v / 2
		counts[string(k[1])] += v / 2
	}

	// ===========================================================================
	// COUNT THE THINGS
	// ===========================================================================

	min := 1000000000000000
	max := 0
	for _, v := range counts {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	return max - min
}
