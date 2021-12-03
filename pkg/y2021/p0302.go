package y2021

import (
	"bufio"
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/timdeklijn/aoc/pkg/solution"
)

const e32 string = `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`

type P0302 struct{}

func (p *P0302) GetExamples() solution.Examples {
	return solution.Examples{
		solution.Example{
			N:    1,
			In:   e32,
			Want: 230,
		},
	}
}

// count the number of '0' and '1' in on position pos of each element of list l.
func count(lines []string, pos int) (int, int) {
	ones := 0
	zeros := 0
	for _, l := range lines {
		switch l[pos] {
		case '0':
			zeros += 1
		case '1':
			ones += 1
		}
	}
	return zeros, ones
}

// filter a list on if the byte in position pos of an element is eqyal to m
func filter(l []string, m byte, pos int) []string {
	tmp := []string{}
	for _, l := range l {
		if l[pos] == m {
			tmp = append(tmp, l)
		}
	}
	return tmp
}

// minMax returns a '1' or '0' depending on if zeros is larger or smaller then
// ones
func minMax(zeros, ones int) (byte, byte) {
	var max byte
	var min byte
	if zeros <= ones {
		max = '1'
		min = '0'
	} else {
		max = '0'
		min = '1'
	}
	return min, max
}

func (p *P0302) Run(data *bufio.Scanner) int {

	lines := []string{}
	for data.Scan() {
		s := data.Text()
		lines = append(lines, s)
	}

	oxygen := lines
	co2 := lines
	ox := true
	co := true

	// for each position in the bit string
	for i := 0; i < len(lines[0]); i++ {

		// zeros and ones per list per posision
		oxZeros, oxOnes := count(oxygen, i)
		coZeros, coOnes := count(co2, i)

		// get the byte to filter on
		_, max := minMax(oxZeros, oxOnes)
		min, _ := minMax(coZeros, coOnes)

		// if oxygen is not of length 1, filter it on the byte
		if ox {
			oxygen = filter(oxygen, max, i)
			if len(oxygen) == 1 {
				ox = false
			}
		}

		// if co2 is not of length 1, filter it on the byte
		if co {
			co2 = filter(co2, min, i)
			if len(co2) == 1 {
				co = false
			}
		}

		// break if both lists are of lenght 1
		if !co && !ox {
			break
		}
	}

	// convert resulting bit-strings to ints
	conum, _ := strconv.ParseInt(co2[0], 2, 64)
	oxnum, _ := strconv.ParseInt(oxygen[0], 2, 64)

	log.Infof("co2: %s - %d", co2, conum)
	log.Infof("oxygen: %s - %d", oxygen, oxnum)

	// return product
	return int(conum * oxnum)
}
