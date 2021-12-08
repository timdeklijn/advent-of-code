package y2021

import (
	"bufio"
	"strings"

	"github.com/timdeklijn/aoc/pkg/solution"
)

const (
	e81 string = `be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe
edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc
fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg
fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb
aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea
fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb
dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe
bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef
egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb
gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce
`
)

type P0801 struct{}

func (p *P0801) GetExamples() solution.Examples {
	return solution.Examples{
		solution.Example{
			N:    1,
			In:   e81,
			Want: 26,
		},
	}
}

func (p *P0801) Run(data *bufio.Scanner) int {

	tot := 0
	for data.Scan() {
		s := data.Text()
		spl := strings.Split(s, "|")
		output := strings.Split(strings.Trim(spl[1], " "), " ")
		for _, i := range output {
			switch len(i) {
			case 2:
				tot++
			case 3:
				tot++
			case 4:
				tot++
			case 7:
				tot++
			default:
				continue
			}
		}
	}

	return tot
}
