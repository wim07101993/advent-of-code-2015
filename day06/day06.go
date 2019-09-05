package day06

import (
	"errors"
	"regexp"
	"strconv"
)

type SolverDay6 struct {
	inputProvider func() []string
	input         []string
	solution1     int
	solution2     int
}

type Grid [][]uint

type Coordinate struct {
	X int
	Y int
}

type Instruction struct {
	From Coordinate
	To   Coordinate
	Do   Action
}

type Action int

const (
	Off Action = iota
	On
	Toggle
)

func New(inputProvider func() []string) *SolverDay6 {
	return &SolverDay6{
		inputProvider: inputProvider,
	}
}

func NewGrid(cols int, rows int) Grid {
	g := make(Grid, cols, cols)
	for x := range g {
		for i := 0; i < rows; i++ {
			g[x] = append(g[x], 0)
		}
	}
	return g
}

func (s *SolverDay6) SolvePart1() string {
	if s.input == nil {
		s.input = s.inputProvider()
	}

	g := NewGrid(1000, 1000)
	for _, in := range s.input {
		i, err := ParseInstruction(in)
		if err != nil {
			panic(err)
		}
		g.Execute(i)
	}
	return strconv.Itoa(g.GetOn())
}

func (s *SolverDay6) SolvePart2() string {
	if s.input == nil {
		s.input = s.inputProvider()
	}

	g := NewGrid(1000, 1000)
	for _, in := range s.input {
		i, err := ParseInstruction(in)
		if err != nil {
			panic(err)
		}
		g.Execute2(i)
	}
	return strconv.Itoa(g.GetOn())
}

func (g Grid) Execute(i Instruction) {
	switch i.Do {
	case Off:
		g.TurnOff(i.From, i.To)
	case On:
		g.TurnOn(i.From, i.To)
	case Toggle:
		g.Toggle(i.From, i.To)
	}
}

func (g Grid) Execute2(i Instruction) {
	switch i.Do {
	case Off:
		g.TurnOff2(i.From, i.To)
	case On:
		g.TurnOn(i.From, i.To)
	case Toggle:
		g.Toggle2(i.From, i.To)
	}
}

func (g Grid) TurnOn(from, to Coordinate) {
	for x := from.X; x <= to.X; x++ {
		for y := from.Y; y <= to.Y; y++ {
			g[x][y]++
		}
	}
}

func (g Grid) TurnOff(from, to Coordinate) {
	for x := from.X; x <= to.X; x++ {
		for y := from.Y; y <= to.Y; y++ {
			g[x][y] = 0
		}
	}
}

func (g Grid) TurnOff2(from, to Coordinate) {
	for x := from.X; x <= to.X; x++ {
		for y := from.Y; y <= to.Y; y++ {
			g[x][y]--
		}
	}
}

func (g Grid) Toggle(from, to Coordinate) {
	for x := from.X; x <= to.X; x++ {
		for y := from.Y; y <= to.Y; y++ {
			if g[x][y] > 0 {
				g[x][y] = 0
			} else {
				g[x][y]++
			}
		}
	}
}

func (g Grid) Toggle2(from, to Coordinate) {
	for x := from.X; x <= to.X; x++ {
		for y := from.Y; y <= to.Y; y++ {
			g[x][y] += 2
		}
	}
}

func (g Grid) GetOn() int {
	i := 0
	for _, col := range g {
		for _, cor := range col {
			if cor > 0 {
				i++
			}
		}
	}
	return i
}

func (g Grid) GetTotalBrightness() uint {
	var b uint = 0
	for _, col := range g {
		for _, cor := range col {
			b += cor
		}
	}
	return b
}

func ParseInstruction(s string) (Instruction, error) {
	i := Instruction{}

	r := regexp.MustCompile(`\d+`)
	ns := r.FindAllString(s, 4)
	if len(ns) != 4 {
		return Instruction{}, errors.New("Could not parse instruction " + s)
	}

	if x, err := strconv.Atoi(ns[0]); err != nil {
		return Instruction{}, err
	} else {
		i.From.X = x
	}

	if y, err := strconv.Atoi(ns[1]); err != nil {
		return Instruction{}, err
	} else {
		i.From.Y = y
	}

	if x, err := strconv.Atoi(ns[2]); err != nil {
		return Instruction{}, err
	} else {
		i.To.X = x
	}

	if y, err := strconv.Atoi(ns[3]); err != nil {
		return Instruction{}, err
	} else {
		i.To.Y = y
	}

	switch s[:7] {
	case "turn of":
		i.Do = Off
	case "turn on":
		i.Do = On
	default:
		i.Do = Toggle
	}

	return i, nil
}
