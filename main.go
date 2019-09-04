package main

import (
	"fmt"

	"github.com/wim07101993/advent-of-code-2015/day01"
	"github.com/wim07101993/advent-of-code-2015/day02"
	"github.com/wim07101993/advent-of-code-2015/day03"
	"github.com/wim07101993/advent-of-code-2015/day04"
	"github.com/wim07101993/advent-of-code-2015/day05"
	"github.com/wim07101993/advent-of-code-2015/utils"
)

var (
	Solvers []utils.Solver
)

func init() {
	Solvers = []utils.Solver{
		day01.New(func() string {
			if s, err := utils.ReadFileAsString("./day01/input"); err != nil {
				panic(err)
			} else {
				return s
			}
		}),
		day02.New(func() []string {
			if s, err := utils.ReadFileLines("./day02/input"); err != nil {
				panic(err)
			} else {
				return s
			}
		}),
		day03.New(func() string {
			if s, err := utils.ReadFileAsString("./day03/input"); err != nil {
				panic(err)
			} else {
				return s
			}
		}),
		day04.New(func() []byte {
			if bs, err := utils.ReadFileAsBytes("./day04/input"); err != nil {
				panic(err)
			} else {
				return bs
			}
		}),
		day05.New(func() []string {
			if bs, err := utils.ReadFileLines("./day05/input"); err != nil {
				panic(err)
			} else {
				return bs
			}
		}),
	}
}

func main() {
	// solve all
	// for i, s := range Solvers {
	// 	solve(s, i+1)
	// }

	solve(Solvers[4], 5)
}

func solve(s utils.Solver, day int) {
	fmt.Println("Day", day, "part 1:", s.SolvePart1())
	fmt.Println("Day", day, "part 2:", s.SolvePart2())
}
