package day04

type SolverDay5 struct {
	inputProvider func() []string
	input         []string
	solution1     int
	solution2     int
}

func New(inputProvider func() []string) *SolverDay5 {
	return &SolverDay5{
		inputProvider: inputProvider,
	}
}

func (s *SolverDay5) SolvePart1() string {
	if s.input == nil {
		s.input = s.inputProvider()
	}

	return ""
}

func (s *SolverDay5) SolvePart2() string {
	if s.input == nil {
		s.input = s.inputProvider()
	}

	return ""
}
