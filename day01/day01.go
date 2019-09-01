package day01

import "strconv"

type SolverDay1 struct {
	input     string
	solution1 int
	solution2 int
}

func New(inputProvider func() string) *SolverDay1 {
	return &SolverDay1{
		input: inputProvider(),
	}
}

func (s *SolverDay1) SolvePart1() string {
	s.solution1 = 0

	for _, c := range s.input {
		if c == '(' {
			s.solution1++
		} else if c == ')' {
			s.solution1--
		}
	}

	return strconv.Itoa(s.solution1)
}

func (s *SolverDay1) SolvePart2() string {
	floor := 0
	s.solution2 = 0

	for i, c := range s.input {
		if c == '(' {
			floor++
		} else if c == ')' {
			floor--
		}

		if floor < 0 {
			s.solution2 = i + 1
			break
		}
	}

	return strconv.Itoa(s.solution2)
}
