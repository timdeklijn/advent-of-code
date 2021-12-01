package solution

import "bufio"

type Example struct {
	N    int
	In   string
	Want int
}

type Examples []Example

type Solution interface {
	GetExamples() Examples
	Run(*bufio.Scanner) int
}
