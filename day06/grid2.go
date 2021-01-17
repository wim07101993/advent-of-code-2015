package day06

import (
	"fmt"
	"strings"
)

type Grid2 [][]int

func NewGrid2(width, height int) Grid2 {
	g := make(Grid2, height, height)
	for y := range g {
		g[y] = make([]int, width, width)
	}
	return g
}

func (g Grid2) Intensity() int {
	var count = 0
	for y := range g {
		for x := range g[y] {
			count+=g[y][x]
		}
	}
	return count
}

func (g Grid2) TurnOn(x, y int) {
	// The phrase turn on actually means that you should increase the brightness
	// of those lights by 1.
	g[y][x] += 1
}

func (g Grid2) TurnOff(x, y int) {
	// The phrase turn off actually means that you should decrease the
	// brightness of those lights by 1, to a minimum of zero.
	g[y][x] -= 1
	if g[y][x] < 0 {
		g[y][x] = 0
	}
}

func (g Grid2) Toggle(x, y int) {
	// The phrase toggle actually means that you should increase the brightness
	// of those lights by 2.
	g[y][x] += 2
}

func (g Grid2) DoForRange(from, to Coordinate, action func(x, y int)) {
	for y := from.Y; y <= to.Y; y++ {
		for x := from.X; x <= to.X; x++ {
			action(x, y)
		}
	}
}

func (g Grid2) String() string {
	b := strings.Builder{}
	for y := range g {
		for x := range g[y] {
			b.WriteString(fmt.Sprintf("%d", g[y][x]))
		}
		b.WriteString("\r\n")
	}
	return b.String()
}
