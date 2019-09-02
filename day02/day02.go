package day02

import (
	"fmt"
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

func New(inputProvider func() []string) *SolverDay2 {
	strInput := inputProvider()
	return &SolverDay2{
		input: ParseInput(strInput),
	}
}

func (s *SolverDay2) SolvePart1() string {
	return ""
}

func (s *SolverDay2) SolvePart2() string {
	return ""
}

func ParseInput(input []string) []WrappingPaperDimensions {
	ret := []WrappingPaperDimensions{}

	l := 0
	w := 0
	h := 0
	var err error
	for _, s := range input {
		split := strings.Split(s, "x")
		if l, err = strconv.Atoi(split[0]); err != nil {
			fmt.Println("Error while parsing input:", err)
			continue
		}
		if w, err = strconv.Atoi(split[1]); err != nil {
			fmt.Println("Error while parsing input:", err)
			continue
		}
		if h, err = strconv.Atoi(split[2]); err != nil {
			fmt.Println("Error while parsing input:", err)
			continue
		}

		ret = append(ret, WrappingPaperDimensions{
			Length: l,
			Width:  w,
			Height: h,
		})
	}

	return ret
}
