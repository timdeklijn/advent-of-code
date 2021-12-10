package y2021

import (
	"bufio"
	"github.com/timdeklijn/aoc/pkg/solution"
)

type P1001 struct{}

const e101 string = `[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]`

func (p *P1001) GetExamples() solution.Examples {
	return solution.Examples{
		solution.Example{
			N:    1,
			In:   e101,
			Want: 26397,
		},
	}
}

func (p *P1001) Run(data *bufio.Scanner) int {

	sum := 0
	points := map[string]int{")": 3, "]": 57, "}": 1197, ">": 25137}

	for data.Scan() {

		s := data.Text()
		var l []string   // fill with opening characters
		var wrong string // the first wrong character
		stop := false    // if true break out of loop for line

		// iterate over characters in line
		for _, i := range s {
			ss := string(i)
			switch ss {
			// opening characters are always added
			case "(", "[", "{", "<":
				l = append(l, ss)

			// if a closing character does not close the last opened character, something
			// is wrong
			case ")":
				if l[len(l)-1] == "(" {
					l = l[:len(l)-1]
				} else {
					wrong = ss
					stop = true
				}

			case "]":
				if l[len(l)-1] == "[" {
					l = l[:len(l)-1]
				} else {
					wrong = ss
					stop = true
				}

			case "}":
				if l[len(l)-1] == "{" {
					l = l[:len(l)-1]
				} else {
					wrong = ss
					stop = true
				}

			case ">":
				if l[len(l)-1] == "<" {
					l = l[:len(l)-1]
				} else {
					wrong = ss
					stop = true
				}

			default:
				continue
			}
			if stop {
				break
			}
		}

		sum += points[wrong]
	}

	return sum
}
