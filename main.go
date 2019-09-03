package main

import (
	"fmt"

	"github.com/wim07101993/advent-of-code-2015/day01"
	"github.com/wim07101993/advent-of-code-2015/day02"
	"github.com/wim07101993/advent-of-code-2015/day03"
	"github.com/wim07101993/advent-of-code-2015/utils"
)

var (
	SolverDay01 *day01.SolverDay1
	SolverDay02 *day02.SolverDay2
	SolverDay03 *day03.SolverDay3
)

func main() {
	initializeSolvers()

	//fmt.Println("Day 01 part 1:", SolverDay01.SolvePart1())
	//fmt.Println("Day 01 part 2:", SolverDay01.SolvePart2())

	// fmt.Println("Day 02 part 1:", SolverDay02.SolvePart1())
	// fmt.Println("Day 02 part 2:", SolverDay02.SolvePart2())

	fmt.Println("Day 03 part 1:", SolverDay03.SolvePart1())
	fmt.Println("Day 03 part 2:", SolverDay03.SolvePart2())
}

func initializeSolvers() {
	SolverDay01 = day01.New(func() string {
		if s, err := utils.ReadFileAsString("./day01/input"); err != nil {
			panic(err)
		} else {
			return s
		}
	})

	SolverDay02 = day02.New(func() []string {
		if s, err := utils.ReadFileLines("./day02/input"); err != nil {
			panic(err)
		} else {
			return s
		}
	})

	SolverDay03 = day03.New(func() string {
		if s, err := utils.ReadFileAsString("./day03/input"); err != nil {
			panic(err)
		} else {
			return s
		}
	})
}
