package day06

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"regexp"
	"strconv"
)

type SolverDay6 struct {
	inputProvider func() []string
	instructions  []Instruction
	solution1     int
	solution2     int
}

type Grid [][]int

type Coordinate struct {
	X int
	Y int
}

type Instruction struct {
	From Coordinate
	To   Coordinate
	Do   Action
}

type Action int

const (
	Off Action = iota
	On
	Toggle
)

func New(inputProvider func() []string) *SolverDay6 {
	return &SolverDay6{
		inputProvider: inputProvider,
	}
}

func NewGrid(cols int, rows int) Grid {
	g := make(Grid, cols, cols)
	for x := range g {
		for i := 0; i < rows; i++ {
			g[x] = append(g[x], 0)
		}
	}
	return g
}

func (s *SolverDay6) SolvePart1() string {
	s.EnsureInstructions()
	g := NewGrid(1000, 1000)
	for _, i := range s.instructions {
		g.Execute(i)
	}
	return strconv.Itoa(g.GetOn())
}

func (s *SolverDay6) SolvePart2() string {
	s.EnsureInstructions()
	g := NewGrid(1000, 1000)
	n := 0

	for _, i := range s.instructions {
		g.Execute2(i)
		f, _ := os.Create("C:/Users/WVA1809/go/src/github.com/wim07101993/advent-of-code-2015/images/" + strconv.Itoa(n) + fmt.Sprint(i) + ".png")
		png.Encode(f, g.ToImage())
		n++
	}
	return strconv.Itoa(g.GetOn())
}

func (s *SolverDay6) EnsureInstructions() {
	if s.instructions == nil {
		input := s.inputProvider()
		for _, in := range input {
			i, err := ParseInstruction(in)
			if err != nil {
				panic(err)
			}
			s.instructions = append(s.instructions, i)
		}
	}
}

func (g Grid) Execute(i Instruction) {
	switch i.Do {
	case Off:
		g.TurnOff(i.From, i.To)
	case On:
		g.TurnOn(i.From, i.To)
	case Toggle:
		g.Toggle(i.From, i.To)
	}
}

func (g Grid) Iterate(from, to Coordinate, action func(value int) int) {
	for x := from.X; x <= to.X; x++ {
		for y := from.Y; y <= to.Y; y++ {
			g[x][y] = action(g[x][y])
		}
	}
}

func (g Grid) TurnOn(from, to Coordinate) {
	g.Iterate(from, to, func(value int) int {
		return value + 1
	})
}

func (g Grid) TurnOff(from, to Coordinate) {
	g.Iterate(from, to, func(value int) int {
		return 0
	})
}

func (g Grid) Toggle(from, to Coordinate) {
	g.Iterate(from, to, func(value int) int {
		if value > 0 {
			return 0
		}
		return 1
	})
}

func (g Grid) Execute2(i Instruction) {
	switch i.Do {
	case Off:
		g.TurnOff2(i.From, i.To)
	case On:
		g.TurnOn(i.From, i.To)
	case Toggle:
		g.Toggle2(i.From, i.To)
	}
}

func (g Grid) TurnOff2(from, to Coordinate) {
	g.Iterate(from, to, func(value int) int {
		value--
		if value < 0 {
			return 0
		}
		return value
	})
}

func (g Grid) Toggle2(from, to Coordinate) {
	g.Iterate(from, to, func(value int) int {
		return value + 2
	})
}

func (g Grid) GetOn() int {
	i := 0
	for _, col := range g {
		for _, cor := range col {
			if cor > 0 {
				i++
			}
		}
	}
	return i
}

func (g Grid) GetTotalBrightness() int {
	var b int = 0
	for _, col := range g {
		for _, cor := range col {
			b += cor
		}
	}
	return b
}

func (g Grid) Print() {
	for _, c := range g {
		fmt.Printf("%v\r\n", c)
	}
}

func (g Grid) ToImage() image.Image {
	upleft := image.Point{0, 0}
	lowRight := image.Point{len(g), len(g[0])}
	img := image.NewRGBA(image.Rectangle{upleft, lowRight})

	for x := range g {
		for y, i := range g[x] {
			v := i * 4
			img.Set(x, y, color.RGBA{
				uint8(v),
				uint8(v),
				uint8(v),
				0xff,
			})
		}
	}

	return img
}

func ParseInstruction(s string) (Instruction, error) {
	i := Instruction{}

	r := regexp.MustCompile(`\d+`)
	ns := r.FindAllString(s, 4)
	if len(ns) != 4 {
		return Instruction{}, errors.New("Could not parse instruction " + s)
	}

	if x, err := strconv.Atoi(ns[0]); err != nil {
		return Instruction{}, err
	} else {
		i.From.X = x
	}

	if y, err := strconv.Atoi(ns[1]); err != nil {
		return Instruction{}, err
	} else {
		i.From.Y = y
	}

	if x, err := strconv.Atoi(ns[2]); err != nil {
		return Instruction{}, err
	} else {
		i.To.X = x
	}

	if y, err := strconv.Atoi(ns[3]); err != nil {
		return Instruction{}, err
	} else {
		i.To.Y = y
	}

	switch s[:7] {
	case "turn of":
		i.Do = Off
	case "turn on":
		i.Do = On
	default:
		i.Do = Toggle
	}

	return i, nil
}
