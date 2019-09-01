package day01

import "strconv"

type SolverPart1 struct {
	input    string
	solution int
}

func NewSolverPart1(inputProvider func() string) *SolverPart1 {
	return &SolverPart1{
		input: inputProvider(),
	}
}

func (s *SolverPart1) Solve() string {
	s.solution = 0

	for _, c := range s.input {
		if c == '(' {
			s.solution++
		} else if c == ')' {
			s.solution--
		}
	}

	return strconv.Itoa(s.solution)
}
