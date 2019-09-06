package main

import (
	"fmt"

	"github.com/wim07101993/advent-of-code-2015/day01"
	"github.com/wim07101993/advent-of-code-2015/day02"
	"github.com/wim07101993/advent-of-code-2015/day03"
	"github.com/wim07101993/advent-of-code-2015/day04"
	"github.com/wim07101993/advent-of-code-2015/day05"
	"github.com/wim07101993/advent-of-code-2015/day06"
	"github.com/wim07101993/advent-of-code-2015/day07"
	"github.com/wim07101993/advent-of-code-2015/day08"
	"github.com/wim07101993/advent-of-code-2015/day09"
	"github.com/wim07101993/advent-of-code-2015/day10"
	"github.com/wim07101993/advent-of-code-2015/day11"
	"github.com/wim07101993/advent-of-code-2015/day12"
	"github.com/wim07101993/advent-of-code-2015/day13"
	"github.com/wim07101993/advent-of-code-2015/day14"
	"github.com/wim07101993/advent-of-code-2015/day15"
	"github.com/wim07101993/advent-of-code-2015/day16"
	"github.com/wim07101993/advent-of-code-2015/day17"
	"github.com/wim07101993/advent-of-code-2015/day18"
	"github.com/wim07101993/advent-of-code-2015/day19"
	"github.com/wim07101993/advent-of-code-2015/day20"
	"github.com/wim07101993/advent-of-code-2015/day21"
	"github.com/wim07101993/advent-of-code-2015/day22"
	"github.com/wim07101993/advent-of-code-2015/day23"
	"github.com/wim07101993/advent-of-code-2015/day24"
	"github.com/wim07101993/advent-of-code-2015/day25"
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
		day10.New(func() string {
			if bs, err := utils.ReadFileAsString("./day10/input"); err != nil {
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

	solve(Solvers[6], 7)
}

func solve(s utils.Solver, day int) {
	fmt.Println("Day", day, "part 1:", s.SolvePart1())
	fmt.Println("Day", day, "part 2:", s.SolvePart2())
}
