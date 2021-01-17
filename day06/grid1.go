package day06

import (
	"fmt"
	"strings"
)

type Grid1 [][]int

func NewGrid1(width, height int) Grid1 {
	g := make(Grid1, height, height)
	for y := range g {
		g[y] = make([]int, width, width)
	}
	return g
}

func (g Grid1) Intensity() int {
	var count = 0
	for y := range g {
		for x := range g[y] {
			if g[y][x] > 0 {
				count+=g[y][x]
			}
		}
	}
	return count
}

func (g Grid1) TurnOn(x, y int) {
	g[y][x] = 1
}

func (g Grid1) TurnOff(x, y int) {
	g[y][x] = 0
}

func (g Grid1) Toggle(x, y int) {
	isOn := g[y][x] > 0
	if isOn {
		g[y][x] = 0
	} else {
		g[y][x] = 1
	}
}

func (g Grid1) DoForRange(from, to Coordinate, action func(x, y int)) {
	for y := from.Y; y <= to.Y; y++ {
		for x := from.X; x <= to.X; x++ {
			action(x, y)
		}
	}
}

func (g Grid1) String() string {
	b := strings.Builder{}
	for y := range g {
		for x := range g[y] {
			b.WriteString(fmt.Sprintf("%d", g[y][x]))
		}
		b.WriteString("\r\n")
	}
	return b.String()
}
