package y2021

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/timdeklijn/aoc/pkg/solution"
)

const e41 string = `7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7`

// =============================================================================
// BINGO TYPES
// =============================================================================

type BingoElement struct {
	num    int
	marked bool
}

// BingoBoard can be indexed like b[row][column]
type BingoBoard map[int]map[int]BingoElement

// Print shows logs the full board (row by row log)
func (b BingoBoard) Print() {
	for x := 0; x < 5; x++ {
		row := []string{}
		for y := 0; y < 5; y++ {
			var n string
			if b[x][y].marked == true {
				n = fmt.Sprintf(" [%d] ", b[x][y].num)
			} else {
				n = fmt.Sprintf("  %d  ", b[x][y].num)
			}
			row = append(row, n)
		}
		log.Infof("   %+v", row)
	}
}

// Mark sets 'marked' for an element with num == i to true
func (b BingoBoard) Mark(i int) {
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			if b[x][y].num == i {
				n := b[x][y].num
				b[x][y] = BingoElement{num: n, marked: true}
			}
		}
	}
}

// CheckRow returns true if a row is completely marked
func (b BingoBoard) CheckRow(r int) bool {
	rr := b[r]
	for i := 0; i < 5; i++ {
		if !rr[i].marked {
			return false
		}
	}
	return true
}

// CheckColumn returns true if a column is completely marked
func (b BingoBoard) CheckColumn(c int) bool {
	for i := 0; i < 5; i++ {
		if !b[i][c].marked {
			return false
		}
	}
	return true
}

// CheckBoard returns true if a row or column is completely marked
func (b BingoBoard) CheckBoard() bool {
	for i := 0; i < 5; i++ {
		if b.CheckRow(i) {
			return true
		}

		if b.CheckColumn(i) {
			return true
		}
	}

	return false
}

// SumUnmarked returns the sum of all unmarked elements on a board
func (b BingoBoard) SumUnmarked() int {
	s := 0
	for r := 0; r < 5; r++ {
		for c := 0; c < 5; c++ {
			if !b[r][c].marked {
				s += b[r][c].num
			}
		}
	}
	return s
}

// NewDraws converts a string of comma separated numbers to a list of ints.
func NewDraws(s string) []int {
	draw := []int{}
	spl := strings.Split(s, ",")
	for _, i := range spl {
		ii, _ := strconv.Atoi(i)
		draw = append(draw, ii)
	}
	return draw
}

// NewBingoBoards converts the boards in the input to []Board
func NewBingoBoards(list []string) []BingoBoard {
	boards := []BingoBoard{}

	b := make(BingoBoard)
	x := 0

	for _, l := range list {
		// If there is an empty line, add the board to the output list and create a
		// new board (new counter as well).
		if l == "" {
			boards = append(boards, b)
			b = make(BingoBoard)
			x = 0
		} else {

			row := make(map[int]BingoElement)
			y := 0

			// This is a bit awkward since the numbers can be 1 or 2 characters long.
			for i := 0; i < 13; i += 3 {
				s := l[i : i+2]
				s = strings.Trim(s, " ")
				num, _ := strconv.Atoi(s)
				row[y] = BingoElement{num: num, marked: false}
				y++
			}

			b[x] = row
			x += 1
		}
	}
	// also include the last board
	boards = append(boards, b)
	return boards
}

// =============================================================================
// SOLUTION
// =============================================================================

type P0401 struct{}

func (p *P0401) GetExamples() solution.Examples {
	return solution.Examples{
		solution.Example{
			N:    1,
			In:   e41,
			Want: 4512,
		},
	}
}

func (p *P0401) Run(data *bufio.Scanner) int {

	// read scanner to list of ints
	lines := []string{}
	for data.Scan() {
		s := data.Text()
		lines = append(lines, s)
	}

	draws := NewDraws(lines[0])
	boards := NewBingoBoards(lines[2:])

	// Loop ever draws, set the numbers as marked in each board and break out of
	// the loop when the first board is finished.
	for _, draw := range draws {
		for _, board := range boards {
			board.Mark(draw)
			if done := board.CheckBoard(); done {
				s := board.SumUnmarked()
				return s * draw
			}
		}
	}

	// Should be unreachable
	return 0
}
