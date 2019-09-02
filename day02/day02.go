package day02

import (
	"strconv"
	"strings"
)

type SolverDay2 struct {
	input     []WrappingPaperDimensions
	solution1 int
	solution2 int
}

type WrappingPaperDimensions struct {
	Length int
	Height int
	Width  int
}

func New(inputProvider func() []string) (*SolverDay2, error) {
	strInput := inputProvider()
	wpds, err := ParseInput(strInput)
	if err != nil {
		return nil, err
	}

	return &SolverDay2{
		input: wpds,
	}, nil
}

func (s *SolverDay2) SolvePart1() string {

	return ""
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
