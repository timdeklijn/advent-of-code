package runner

import (
	"fmt"
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
	// TODO: add data download and loader.
	dl := dataloader.NewDataLoader(r.Day, r.Year)
	data, err := dl.FetchData()
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(string(data))

	sol := p.RunSolution()
	log.Infof("solution: %d", sol)
}
