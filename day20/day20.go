package day20

type SolverDay20 struct {
	input     string
	solution1 int
	solution2 int
}

func New(inputProvider func() string) *SolverDay20 {
	return &SolverDay20{
		input: inputProvider(),
	}
}

func (s *SolverDay20) SolvePart1() string {
	return ""
}

func (s *SolverDay20) SolvePart2() string {
	return ""
}
