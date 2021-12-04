package y2021

import (
	"bufio"

	"github.com/timdeklijn/aoc/pkg/solution"
)

type P0402 struct{}

func (p *P0402) GetExamples() solution.Examples {
	return solution.Examples{
		solution.Example{
			N:    1,
			In:   e41,
			Want: 1924,
		},
	}
}

func (p *P0402) Run(data *bufio.Scanner) int {

	// read scanner to list of ints
	lines := []string{}
	for data.Scan() {
		s := data.Text()
		lines = append(lines, s)
	}

	// Parse dhe drawn numbers as well as the boards
	draws := NewDraws(lines[0])
	boards := NewBingoBoards(lines[2:])

	for _, draw := range draws {

		// Mark all boards
		for _, board := range boards {
			board.Mark(draw)
		}

		// If the final board has a full row or column it should be returned.
		if len(boards) == 1 && boards[0].CheckBoard() {
			s := boards[0].SumUnmarked()
			return s * draw
		}

		// remove all completed boards from 'boards'
		tmpBoards := []BingoBoard{}
		for _, board := range boards {
			if done := board.CheckBoard(); !done {
				tmpBoards = append(tmpBoards, board)
			}
		}
		boards = tmpBoards
	}

	// Should be unreachable
	return 0
}
