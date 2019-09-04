package day10

type SolverDay10 struct {
	input     string
	solution1 int
	solution2 int
}

func New(inputProvider func() string) *SolverDay10 {
	return &SolverDay10{
		input: inputProvider(),
	}
}

func (s *SolverDay10) SolvePart1() string {
	return ""
}

func (s *SolverDay10) SolvePart2() string {
	return ""
}
