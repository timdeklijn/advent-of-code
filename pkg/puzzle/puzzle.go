package puzzle

import (
	log "github.com/sirupsen/logrus"
	"github.com/timdeklijn/aoc/pkg/solution"
)

type Puzzle struct{}

// RunExamples will run the examples through the solution of the specific
// puzzle. When these tests fail false will be returned. If tests pass 'true' is
// returend and the process will be continued.
func (p *Puzzle) RunExamples(s solution.Solution) bool {
	for _, e := range s.GetExamples() {
		got := s.Run(e.In)
		if got != e.Want {
			log.Warningf("example %d failed. Got: %d, want: %d", e.N, got, e.Want)
			return false
		}
		log.Infof("example %d passed", e.N)
	}
	return true
}

func (p *Puzzle) RunSolution() int {
	return 1
}
