package day08

import (
	"fmt"
	"strings"
)

type SolverDay8 struct {
	inputProvider func() []string
	input         []string
	solution1     int
	solution2     int
}

func New(inputProvider func() []string) *SolverDay8 {
	return &SolverDay8{
		inputProvider: inputProvider,
	}
}

func (s *SolverDay8) SolvePart1() string {
	if s.input == nil {
		s.input = s.inputProvider()
	}

	return ""
}

func (s *SolverDay8) SolvePart2() string {
	if s.input == nil {
		s.input = s.inputProvider()
	}

	return ""
}

func CompileString(s string) string {
	if !strings.Contains(s, "\\") {
		return strings.Replace(s, "\"", "", -1)
	}

	s = s[1 : len(s)-1]
	fmt.Println(s)
	return s
}
