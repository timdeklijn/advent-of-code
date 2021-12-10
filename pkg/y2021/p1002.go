package y2021

import (
	"bufio"
	"github.com/timdeklijn/aoc/pkg/solution"
	"sort"
)

type P1002 struct{}

func (p *P1002) GetExamples() solution.Examples {
	return solution.Examples{
		solution.Example{
			N:    1,
			In:   e101,
			Want: 288957,
		},
	}
}

func (p *P1002) Run(data *bufio.Scanner) int {

	var scoreList []int
	points := map[string]int{"(": 1, "[": 2, "{": 3, "<": 4}

	for data.Scan() {

		s := data.Text()
		var l []string // fill with opening characters
		stop := false  // if true break out of loop for line

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
					stop = true
				}

			case "]":
				if l[len(l)-1] == "[" {
					l = l[:len(l)-1]
				} else {
					stop = true
				}

			case "}":
				if l[len(l)-1] == "{" {
					l = l[:len(l)-1]
				} else {
					stop = true
				}

			case ">":
				if l[len(l)-1] == "<" {
					l = l[:len(l)-1]
				} else {
					stop = true
				}

			default:
				continue
			}
			if stop {
				break
			}
		}

		// if we set stop it means the line is corrupted. Corrupted lines are not taken
		// into account to calculate the autocomplete string.
		if stop {
			continue
		}

		// Calculate the score for the closing string based on the remaining open
		// characters in the leftover slice.
		score := 0
		for i := len(l) - 1; i >= 0; i-- {
			score = score*5 + points[l[i]]
		}
		scoreList = append(scoreList, score)
	}

	// Sort the list and return the middle element of the list as answer.
	sort.Ints(scoreList)
	return scoreList[len(scoreList)/2]
}
