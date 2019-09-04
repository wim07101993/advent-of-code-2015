package day21

type SolverDay21 struct {
	inputProvider func() []string
	input         []string
	solution1     int
	solution2     int
}

func New(inputProvider func() []string) *SolverDay21 {
	return &SolverDay21{
		inputProvider: inputProvider,
	}
}

func (s *SolverDay21) SolvePart1() string {
	if s.input == nil {
		s.input = s.inputProvider()
	}

	return ""
}

func (s *SolverDay21) SolvePart2() string {
	if s.input == nil {
		s.input = s.inputProvider()
	}

	return ""
}
