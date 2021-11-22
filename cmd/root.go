package cmd

import (
	"time"

	"github.com/spf13/cobra"
	"github.com/timdeklijn/aoc/pkg/runner"
)

var year int
var day int
var part int

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "aoc",
	Short: "Advent of Code Solutions",
	Long: `A go package simplifying the running of solutions for Advent of
Code. Automate running the examples before downloading the data and running
the solution on the acutal data. Time the solutions and save the run time and
solution to a file.`,

	Run: func(cmd *cobra.Command, args []string) {
		aocRunner := runner.NewRunner(year, day, part)
		aocRunner.Run()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

// setDefaults returns a PuzzleID that, when it is December returns the current
// year + day of the month. If not, it will return the default 2021, 1, 1 which
// means the first part of the first puzzle of 2020.
func setDefaults() (int, int, int) {
	currentTime := time.Now()
	if currentTime.Month() != time.December {
		return 2020, 1, 1
	}
	return currentTime.Year(), currentTime.Day(), 1
}

func init() {
	y, d, p := setDefaults()
	// TODO: this is not set properly
	rootCmd.PersistentFlags().IntVarP(&year, "year", "y", y, "year of puzzle")
	rootCmd.PersistentFlags().IntVarP(&day, "day", "d", d, "day of puzzle")
	rootCmd.PersistentFlags().IntVarP(&part, "part", "p", p, "part of puzzle")
}
