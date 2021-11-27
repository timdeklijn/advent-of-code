package runner

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/timdeklijn/aoc/pkg/dataloader"
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
	solution, err := puzzle.Map(r.Year, r.Day, r.Part)
	if err != nil {
		log.Panic(err)
	}
	p := puzzle.Puzzle{}

	// Run the examples, if this fails, exit the program
	if !p.RunExamples(solution) {
		log.Error("Examples failed")
		os.Exit(0)
	}
	log.Info("examples succeeded 😎")

	dl, err := dataloader.NewDataLoader(r.Day, r.Year)

	if err != nil {
		log.Panicf("error creating dataloader: %e", err)
	}
	scanner, err := dl.RetrieveData()
	if err != nil {
		log.Panicf("error retrieving data: %e", err)
	}

	sol := p.RunSolution(scanner, solution)
	log.Infof("solution: %d", sol)
}
