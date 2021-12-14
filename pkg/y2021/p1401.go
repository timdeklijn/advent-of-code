package y2021

import (
	"bufio"
	"strings"

	"github.com/timdeklijn/aoc/pkg/solution"
)

type P1401 struct{}

const e141 string = `NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C`

func (p *P1401) GetExamples() solution.Examples {
	return solution.Examples{
		solution.Example{
			N:    1,
			In:   e141,
			Want: 1588,
		},
	}
}

func applyRules(template string, rules map[string]string) string {
	new := string(template[0])
	for i := 1; i < len(template); i++ {
		pair := string(new[len(new)-1]) + string(template[i])
		new += rules[pair] + string(template[i])
	}
	return new
}

func countOccurences(s string) map[string]int {
	count := map[string]int{}
	for _, i := range s {
		count[string(i)] += 1
	}
	return count
}

func (p *P1401) Run(data *bufio.Scanner) int {

	var template string
	rules := map[string]string{}
	temp := true
	for data.Scan() {
		s := data.Text()
		if s == "" {
			temp = false
			continue
		}

		if temp {
			template = s
			continue
		}

		spl := strings.Split(s, " -> ")
		rules[spl[0]] = spl[1]
	}

	for i := 0; i < 10; i++ {
		template = applyRules(template, rules)
	}

	count := countOccurences(template)

	min := 1000000
	max := 0
	for _, v := range count {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	return max - min
}
