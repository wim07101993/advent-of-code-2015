package day11

type SolverDay11 struct {
	input     string
	solution1 int
	solution2 int
}

func New(inputProvider func() string) *SolverDay11 {
	return &SolverDay11{
		input: inputProvider(),
	}
}

func (s *SolverDay11) SolvePart1() string {
	p := NewPassword(s.input)
	p.Next()
	return p.String()
}

func (s *SolverDay11) SolvePart2() string {
	return ""
}
