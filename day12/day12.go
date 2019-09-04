package day12

type SolverDay12 struct {
	inputProvider func() string
	input         string
	solution1     int
	solution2     int
}

func New(inputProvider func() string) *SolverDay12 {
	return &SolverDay12{
		inputProvider: inputProvider,
	}
}

func (s *SolverDay12) SolvePart1() string {
	if s.input == "" {
		s.input = s.inputProvider()
	}

	return ""
}

func (s *SolverDay12) SolvePart2() string {
	if s.input == "" {
		s.input = s.inputProvider()
	}

	return ""
}
