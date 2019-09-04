package day07

type SolverDay7 struct {
	inputProvider func() []string
	input         []string
	solution1     int
	solution2     int
}

func New(inputProvider func() []string) *SolverDay7 {
	return &SolverDay7{
		inputProvider: inputProvider,
	}
}

func (s *SolverDay7) SolvePart1() string {
	if s.input == nil {
		s.input = s.inputProvider()
	}

	return ""
}

func (s *SolverDay7) SolvePart2() string {
	if s.input == nil {
		s.input = s.inputProvider()
	}

	return ""
}
