package runner

import (
	log "github.com/sirupsen/logrus"
	"github.com/timdeklijn/aoc/pkg/puzzle"
)

type Runner struct {
	Year, Day, Part int
}

func NewRunner(y, d, p int) Runner {
	return Runner{Year: y, Day: d, Part: p}
}

func (r *Runner) Run() {
	log.Infof("Running puzzle year: %d, Day: %d, Part: %d", r.Year, r.Day, r.Part)
	solution := puzzle.Map(r.Year, r.Day, r.Part)
	p := puzzle.Puzzle{}
	if !p.RunExamples(solution) {
		log.Warning("Examples failed")
	}
	// TODO: add data download and loader.
	sol := p.RunSolution()
	log.Infof("solution: %d", sol)
}
