package puzzle

import (
	"bufio"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/timdeklijn/aoc/pkg/solution"
)

type Puzzle struct{}

func toScanner(s string) *bufio.Scanner {
	return bufio.NewScanner(strings.NewReader(s))
}

// RunExamples will run the examples through the solution of the specific
// puzzle. When these tests fail false will be returned. If tests pass 'true' is
// returend and the process will be continued.
func (p *Puzzle) RunExamples(s solution.Solution) bool {
	for _, e := range s.GetExamples() {
		got := s.Run(toScanner(e.In))
		if got != e.Want {
			log.Warningf("example %d failed. Got: %d, want: %d", e.N, got, e.Want)
			return false
		}
		log.Infof("example %d passed", e.N)
	}
	return true
}

func (p *Puzzle) RunSolution(data *bufio.Scanner, s solution.Solution) int {
	return s.Run(data)
}
