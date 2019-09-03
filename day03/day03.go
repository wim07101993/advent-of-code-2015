package day03

import "strconv"

const (
	santa = 0
	robot = 1
)

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
	return strconv.Itoa(calcDeliveredHouses(s.input))
}

func (s *SolverDay3) SolvePart2() string {
	return strconv.Itoa(calcDeliveredHousesWithRobot(s.input))
}

func calcDeliveredHouses(input string) int {
	sc := coordinate{}

	deliveries := map[coordinate]int{
		sc: 1,
	}

	for _, m := range input {
		increaseCoordinate(m, &sc)
		if _, ok := deliveries[sc]; !ok {
			deliveries[sc]++
		}
	}

	return len(deliveries)
}

func calcDeliveredHousesWithRobot(input string) int {
	sc := coordinate{}
	rc := coordinate{}
	turn := santa

	deliveries := map[coordinate]int{
		sc: 2,
	}

	for _, m := range input {
		if turn == santa {
			increaseCoordinate(m, &sc)
			if _, ok := deliveries[sc]; !ok {
				deliveries[sc]++
			}
			turn = robot
		} else {
			increaseCoordinate(m, &rc)
			if _, ok := deliveries[rc]; !ok {
				deliveries[rc]++
			}
			turn = santa
		}
	}

	return len(deliveries)
}

func increaseCoordinate(r rune, c *coordinate) {
	switch r {
	case '^':
		c.Y++
	case 'v':
		c.Y--
	case '>':
		c.X++
	case '<':
		c.X--
	}
}
