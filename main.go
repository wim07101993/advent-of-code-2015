package main

import (
	"fmt"

	"github.com/wim07101993/advent-of-code-2015/day01"
	"github.com/wim07101993/advent-of-code-2015/utils"
)

var (
	SolverDay011 *day01.SolverPart1
)

func main() {
	initializeSolvers()

	fmt.Println("Day 01 part 1:", SolverDay011.Solve())
}

func initializeSolvers() {
	SolverDay011 = day01.NewSolverPart1(func() string {
		if s, err := utils.ReadFileAsString("./day01/input"); err != nil {
			panic(err)
		} else {
			return s
		}
	})

}
