package day06

import (
	"strconv"
)

type SolverDay6 struct {
	inputProvider func() []string
	instructions  []Instruction
	solution1     int
	solution2     int
	grid Grid
}

func New(inputProvider func() []string) *SolverDay6 {
	return &SolverDay6{
		inputProvider: inputProvider,
	}
}

func (s *SolverDay6) SolvePart1() string {
	s.EnsureInstructions()
	s.grid = NewGrid1(1000, 1000)
	for _, i := range s.instructions {
		i.ExecuteOn(s.grid)
	}
	return strconv.Itoa(s.grid.Intensity())
}

func (s *SolverDay6) SolvePart2() string {
	s.EnsureInstructions()
	s.grid = NewGrid2(1000, 1000)

	for _, i := range s.instructions {
		i.ExecuteOn(s.grid)
	}
	return strconv.Itoa(s.grid.Intensity())
}

func (s *SolverDay6) EnsureInstructions() {
	if s.instructions == nil {
		input := s.inputProvider()
		for _, in := range input {
			i, err := ParseInstruction(in)
			if err != nil {
				panic(err)
			}
			s.instructions = append(s.instructions, i)
		}
	}
}
