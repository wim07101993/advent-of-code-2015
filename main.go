package main

import (
	"advent-of-code-2015/day01"
	"advent-of-code-2015/day02"
	"advent-of-code-2015/day03"
	"advent-of-code-2015/day04"
	"advent-of-code-2015/day05"
	"advent-of-code-2015/day06"
	"advent-of-code-2015/day07"
	"advent-of-code-2015/day08"
	"advent-of-code-2015/day09"
	"advent-of-code-2015/day10"
	"advent-of-code-2015/day11"
	"advent-of-code-2015/day12"
	"advent-of-code-2015/day13"
	"advent-of-code-2015/day14"
	"advent-of-code-2015/day15"
	"advent-of-code-2015/day16"
	"advent-of-code-2015/day17"
	"advent-of-code-2015/day18"
	"advent-of-code-2015/day19"
	"advent-of-code-2015/day20"
	"advent-of-code-2015/day21"
	"advent-of-code-2015/day22"
	"advent-of-code-2015/day23"
	"advent-of-code-2015/day24"
	"advent-of-code-2015/day25"
	"advent-of-code-2015/utils"
	"fmt"

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
		day06.New(func() []string {
			if bs, err := utils.ReadFileLines("./day06/input"); err != nil {
				panic(err)
			} else {
				return bs
			}
		}),
		day07.New(func() []string {
			if bs, err := utils.ReadFileLines("./day07/input"); err != nil {
				panic(err)
			} else {
				return bs
			}
		}),
		day08.New(func() []string {
			if bs, err := utils.ReadFileLines("./day08/input"); err != nil {
				panic(err)
			} else {
				return bs
			}
		}),
		day09.New(func() []string {
			if bs, err := utils.ReadFileLines("./day09/input"); err != nil {
				panic(err)
			} else {
				return bs
			}
		}),
		day10.New(func() []byte {
			if bs, err := utils.ReadFileAsBytes("./day10/input"); err != nil {
				panic(err)
			} else {
				return bs
			}
		}),
		day11.New(func() string {
			if bs, err := utils.ReadFileAsString("./day11/input"); err != nil {
				panic(err)
			} else {
				return bs
			}
		}),
		day12.New(func() string {
			if bs, err := utils.ReadFileAsString("./day12/input"); err != nil {
				panic(err)
			} else {
				return bs
			}
		}),
		day13.New(func() []string {
			if bs, err := utils.ReadFileLines("./day13/input"); err != nil {
				panic(err)
			} else {
				return bs
			}
		}),
		day14.New(func() []string {
			if bs, err := utils.ReadFileLines("./day14/input"); err != nil {
				panic(err)
			} else {
				return bs
			}
		}),
		day15.New(func() []string {
			if bs, err := utils.ReadFileLines("./day15/input"); err != nil {
				panic(err)
			} else {
				return bs
			}
		}),
		day16.New(func() []string {
			if bs, err := utils.ReadFileLines("./day16/input"); err != nil {
				panic(err)
			} else {
				return bs
			}
		}),
		day17.New(func() []string {
			if bs, err := utils.ReadFileLines("./day17/input"); err != nil {
				panic(err)
			} else {
				return bs
			}
		}),
		day18.New(func() []string {
			if bs, err := utils.ReadFileLines("./day/18/input"); err != nil {
				panic(err)
			} else {
				return bs
			}
		}),
		day19.New(func() []string {
			if bs, err := utils.ReadFileLines("./day19/input"); err != nil {
				panic(err)
			} else {
				return bs
			}
		}),
		day20.New(func() string {
			if bs, err := utils.ReadFileAsString("./day20/input"); err != nil {
				panic(err)
			} else {
				return bs
			}
		}),
		day21.New(func() []string {
			if bs, err := utils.ReadFileLines("./day21/input"); err != nil {
				panic(err)
			} else {
				return bs
			}
		}),
		day22.New(func() []string {
			if bs, err := utils.ReadFileLines("./day22/input"); err != nil {
				panic(err)
			} else {
				return bs
			}
		}),
		day23.New(func() []string {
			if bs, err := utils.ReadFileLines("./day23/input"); err != nil {
				panic(err)
			} else {
				return bs
			}
		}),
		day24.New(func() []string {
			if bs, err := utils.ReadFileLines("./day24/input"); err != nil {
				panic(err)
			} else {
				return bs
			}
		}),
		day25.New(func() []string {
			if bs, err := utils.ReadFileLines("./day25/input"); err != nil {
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

	solve(Solvers[11], 12)
}

func solve(s utils.Solver, day int) {
	fmt.Println("Day", day, "part 1:", s.SolvePart1())
	fmt.Println("Day", day, "part 2:", s.SolvePart2())
}
