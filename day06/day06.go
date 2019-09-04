package day06

type SolverDay6 struct {
	inputProvider func() []string
	input         []string
	solution1     int
	solution2     int
}

type Grid [1000][1000]bool

type Coordinate struct {
	X int
	Y int
}

func New(inputProvider func() []string) *SolverDay6 {
	return &SolverDay6{
		inputProvider: inputProvider,
	}
}

func (s *SolverDay6) SolvePart1() string {
	if s.input == nil {
		s.input = s.inputProvider()
	}

	return ""
}

func (s *SolverDay6) SolvePart2() string {
	if s.input == nil {
		s.input = s.inputProvider()
	}

	return ""
}

func (g *Grid) Set(from, to Coordinate, value bool) {
	for x := from.X; x < to.X; x++ {
		for y := from.Y; y < to.Y; y++ {
			g[x][y] = value
		}
	}
}

func (g *Grid) Toggle(from, to Coordinate) {
	for x := from.X; x < to.X; x++ {
		for y := from.Y; y < to.Y; y++ {
			g[x][y] = !g[x][y]
		}
	}
}

func (g *Grid) GetOn() int {
	i := 0
	for _, col := range g {
		for _, cor := range col {
			if cor {
				i++
			}
		}
	}
	return i
}
