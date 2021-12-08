package y2021

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/timdeklijn/aoc/pkg/solution"
)

const (
	e82 string = `acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf `
	e83 string = `be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe`
)

type P0802 struct{}

func (p *P0802) GetExamples() solution.Examples {
	return solution.Examples{
		solution.Example{
			N:    1,
			In:   e82,
			Want: 5353,
		},
		solution.Example{
			N:    2,
			In:   e83,
			Want: 8394,
		},
	}
}

// =============================================================================
// Number helps manupulation
// =============================================================================

type Number string

// newNumber returns a Number without duplicates and sorted
func newNumber(s string) Number {
	return Number(s).unique().sort()
}

// sort returns the characters in Number sorted
func (n Number) sort() Number {
	w := strings.Split(string(n), "")
	sort.Strings(w)
	return Number(strings.Join(w, ""))
}

// unique returns the Number without internal duplicates
func (n Number) unique() Number {

	l := strings.Split(string(n), "")
	keys := make(map[string]bool)

	var list []string
	for _, entry := range l {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return Number(strings.Join(list, ""))
}

// add returns the two numbers added, ignoring the duplicates and sorting the
// result
func (n Number) add(o Number) Number {
	return (n + o).unique().sort()
}

// sub returns the characters in o removed from n, sorted.
func (n Number) sub(o Number) Number {
	out := ""
	for _, nn := range string(n) {
		found := false
		for _, oo := range string(o) {
			if nn == oo {
				found = true
			}
		}
		if !found {
			out += string(nn)
		}
	}
	return Number(out).sort()
}

func (p *P0802) Run(data *bufio.Scanner) int {

	tot := 0
	for data.Scan() {
		s := data.Text()
		spl := strings.Split(s, "|")

		// place inputs in a map with length as key
		input := strings.Split(strings.Trim(spl[0], " "), " ")
		inputs := map[int][]string{}
		for _, s := range input {
			inputs[len(s)] = append(inputs[len(s)], s)
		}

		output := strings.Split(strings.Trim(spl[1], " "), " ")

		// =====================================================================
		// Figure out what code belongs to what number
		// =====================================================================

		results := map[int]Number{}

		// these are based on length
		results[1] = newNumber(inputs[2][0])
		results[4] = newNumber(inputs[4][0])
		results[7] = newNumber(inputs[3][0])
		results[8] = newNumber(inputs[7][0])

		// Adding 1 + 7 + 4 results in something with a difference of two to
		// one of the 5 long inputs
		sumNumber := results[1].add(results[7]).add(results[4])
		for _, num := range inputs[5] {
			n := newNumber(num)
			if len(n.sub(sumNumber)) == 2 {
				results[2] = n
			}
		}

		// If we know '2' we can identify 3 and 5 based on the number of
		// different characters
		for _, num := range inputs[5] {
			n := newNumber(num)
			switch len(n.sub(results[2])) {
			case 1:
				results[3] = n
			case 2:
				results[5] = n
			default:
				continue
			}
		}

		// Identify the 6 long inputs
		for _, num := range inputs[6] {
			n := newNumber(num)
			// '9' is the same as '4' + '5'
			if n == results[4].add(results[5]) {
				results[9] = n
				continue
			}
			// '0' will have total overlap with 0
			if len(results[1].sub(n)) == 0 {
				results[0] = n
			} else {
				// '6' misses the top part of 1
				results[6] = n
			}
		}

		// =====================================================================
		// Got a mapping, now create output
		// =====================================================================

		// invert the map to get the results
		trans := map[Number]int{}
		for k, v := range results {
			trans[v] = k
		}

		// get the 4 numbers and combine to a large single numbers
		out := ""
		for _, o := range output {
			n := newNumber(o)
			out += fmt.Sprintf("%d", trans[n])
		}

		// Add the single large number to the total
		subTot, _ := strconv.Atoi(out)
		tot += subTot
	}

	return tot
}
