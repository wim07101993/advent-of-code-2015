package day18

type SolverDay18 struct {
	inputProvider func() []string
	input         []string
	solution1     int
	solution2     int
}

func New(inputProvider func() []string) *SolverDay18 {
	return &SolverDay18{
		inputProvider: inputProvider,
	}
}

func (s *SolverDay18) SolvePart1() string {
	if s.input == nil {
		s.input = s.inputProvider()
	}

	return ""
}

func (s *SolverDay18) SolvePart2() string {
	if s.input == nil {
		s.input = s.inputProvider()
	}

	return ""
}
