package day02

import (
	"strconv"
	"strings"

	"advent-of-code-2015/utils"
)

type SolverDay2 struct {
	inputProvider func() []string
	input         []Present
	solution1     int
	solution2     int
}

type Present struct {
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
	ps, err := ParseInput(strInput)
	if err != nil {
		return err
	}
	s.input = ps
	return nil
}

func (s *SolverDay2) SolvePart1() string {
	if s.input == nil {
		s.fetchInput()
	}
	total := 0
	for _, p := range s.input {
		total += p.CalcWrappingPaperArea()
	}
	return strconv.Itoa(total)
}

func (s *SolverDay2) SolvePart2() string {
	if s.input == nil {
		s.fetchInput()
	}

	total := 0
	for _, p := range s.input {
		total += p.CalcRibbonLength()
	}
	return strconv.Itoa(total)
}

func ParseInput(input []string) ([]Present, error) {
	ret := []Present{}
	for _, s := range input {
		p, err := ParsePresent(s)
		if err != nil {
			return nil, err
		}
		ret = append(ret, p)
	}
	return ret, nil
}

func ParsePresent(input string) (Present, error) {
	l := 0
	w := 0
	h := 0
	var err error

	split := strings.Split(input, "x")
	if l, err = strconv.Atoi(split[0]); err != nil {
		return Present{}, err
	}
	if w, err = strconv.Atoi(split[1]); err != nil {
		return Present{}, err
	}
	if h, err = strconv.Atoi(split[2]); err != nil {
		return Present{}, err
	}

	return Present{
		Length: l,
		Width:  w,
		Height: h,
	}, nil
}

func (p Present) CalcWrappingPaperArea() int {
	areas := []int{
		p.Length * p.Width,
		p.Width * p.Height,
		p.Height * p.Length,
	}

	min := utils.Min(areas...)
	return utils.Sum(areas...)*2 + min
}

func (p Present) CalcRibbonLength() int {
	max := utils.Max(p.Width, p.Length, p.Height)

	if max == p.Width {
		return 2*p.Length + 2*p.Height + p.Length*p.Width*p.Height
	}
	if max == p.Length {
		return 2*p.Width + 2*p.Height + p.Length*p.Width*p.Height
	}

	return 2*p.Width + 2*p.Length + p.Length*p.Width*p.Height
}
