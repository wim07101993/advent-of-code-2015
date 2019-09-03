package day03

import "strconv"

type SolverDay3 struct {
	input     string
	solution1 int
	solution2 int
}

type coordinate struct {
	X int
	Y int
}

func New(inputProvider func() string) *SolverDay3 {
	return &SolverDay3{
		input: inputProvider(),
	}
}

func (s *SolverDay3) SolvePart1() string {
	return strconv.Itoa(calcNumberOfDeliveredHouses(s.input))
}

func (s *SolverDay3) SolvePart2() string {
	return ""
}

func calcNumberOfDeliveredHouses(input string) int {
	cc := coordinate{}

	deliveries := map[coordinate]int{
		cc: 1,
	}

	for _, m := range input {
		switch m {
		case '^':
			cc.Y++
		case 'v':
			cc.Y--
		case '>':
			cc.X++
		case '<':
			cc.X--
		}

		deliveries[cc] = deliveries[cc] + 1
	}

	return len(deliveries)
}
