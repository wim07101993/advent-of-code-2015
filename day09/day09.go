package day09

import (
	"log"
	"strconv"
)

type SolverDay9 struct {
	inputProvider func() []string
	input         []string
	solution1     int
	solution2     int
}

func New(inputProvider func() []string) *SolverDay9 {
	return &SolverDay9{
		inputProvider: inputProvider,
	}
}

func (s *SolverDay9) SolvePart1() string {
	if s.input == nil {
		s.input = s.inputProvider()
	}

	m := ParseMap(s.input)
	d, h := m.ShortestDistance()

	log.Println(h)
	return strconv.Itoa(d)
}

func (s *SolverDay9) SolvePart2() string {
	if s.input == nil {
		s.input = s.inputProvider()
	}
	return ""
}
