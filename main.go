package main

import (
	"fmt"

	"github.com/wim07101993/advent-of-code-2015/day01"
	"github.com/wim07101993/advent-of-code-2015/utils"
)

var (
	SolverDay01 *day01.SolverDay1
)

func main() {
	initializeSolvers()

	fmt.Println("Day 01 part 1:", SolverDay01.SolvePart1())
	fmt.Println("Day 01 part 2:", SolverDay01.SolvePart2())
}

func initializeSolvers() {
	SolverDay01 = day01.New(func() string {
		if s, err := utils.ReadFileAsString("./day01/input"); err != nil {
			panic(err)
		} else {
			return s
		}
	})

}
