package day23

type SolverDay23 struct {
	inputProvider func() []string
	input         []string
	solution1     int
	solution2     int
}

func New(inputProvider func() []string) *SolverDay23 {
	return &SolverDay23{
		inputProvider: inputProvider,
	}
}

func (s *SolverDay23) SolvePart1() string {
	if s.input == nil {
		s.input = s.inputProvider()
	}

	return ""
}

func (s *SolverDay23) SolvePart2() string {
	if s.input == nil {
		s.input = s.inputProvider()
	}

	return ""
}
