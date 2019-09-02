package day02

import (
	"strconv"
	"strings"
)

type SolverDay2 struct {
	inputProvider func() []string
	input         []WrappingPaperDimensions
	solution1     int
	solution2     int
}

type WrappingPaperDimensions struct {
	Length int
	Height int
	Width  int
}

func New(inputProvider func() []string) *SolverDay2 {
	return &SolverDay2{
		inputProvider: inputProvider,
	}
}

func (s *SolverDay2) fetchInput() error {
	strInput := s.inputProvider()
	wpds, err := ParseInput(strInput)
	if err != nil {
		return err
	}
	s.input = wpds
	return nil
}

func (s *SolverDay2) SolvePart1() string {
	if s.input == nil {
		s.fetchInput()
	}
	total := 0
	for _, wpd := range s.input {
		total += wpd.CalcArea()
	}
	return strconv.Itoa(total)
}

func (s *SolverDay2) SolvePart2() string {
	return ""
}

func ParseInput(input []string) ([]WrappingPaperDimensions, error) {
	ret := []WrappingPaperDimensions{}
	for _, s := range input {
		wpd, err := ParseWrappingPaperDimensions(s)
		if err != nil {
			return nil, err
		}
		ret = append(ret, wpd)
	}
	return ret, nil
}

func ParseWrappingPaperDimensions(input string) (WrappingPaperDimensions, error) {
	l := 0
	w := 0
	h := 0
	var err error

	split := strings.Split(input, "x")
	if l, err = strconv.Atoi(split[0]); err != nil {
		return WrappingPaperDimensions{}, err
	}
	if w, err = strconv.Atoi(split[1]); err != nil {
		return WrappingPaperDimensions{}, err
	}
	if h, err = strconv.Atoi(split[2]); err != nil {
		return WrappingPaperDimensions{}, err
	}

	return WrappingPaperDimensions{
		Length: l,
		Width:  w,
		Height: h,
	}, nil
}

func (w WrappingPaperDimensions) CalcArea() int {
	area := w.Length * w.Width
	min := area
	areas := 2 * area

	area = w.Width * w.Height
	areas += 2 * area
	if area < min {
		min = area
	}

	area = w.Height * w.Length
	areas += 2 * area
	if area < min {
		min = area
	}

	return areas + min
}
