package day10

import (
	"strconv"
)

type SolverDay10 struct {
	input     []byte
	solution1 int
	solution2 int
}

func New(inputProvider func() []byte) *SolverDay10 {
	return &SolverDay10{
		input: inputProvider(),
	}
}

func (s *SolverDay10) SolvePart1() string {
	bs := Sequence(s.input, 40)
	return strconv.Itoa(len(bs))
}

func (s *SolverDay10) SolvePart2() string {
	bs := Sequence(s.input, 50)
	return strconv.Itoa(len(bs))
}

func Sequence(bs []byte, count int) []byte {
	for i := 0; i < count; i++ {
		bs = Increase(bs)
	}
	return bs
}

func Increase(bs []byte) []byte {
	if len(bs) == 1 {
		return append([]byte{'1'}, bs...)
	}

	ret := []byte{}
	last := bs[0]
	count := 1

	for i := 1; i < len(bs); i++ {
		if bs[i] == last {
			count++
		} else {
			sCount := strconv.Itoa(count)
			ret = append(ret, []byte(sCount)...)
			ret = append(ret, last)
			last = bs[i]
			count = 1
		}
	}

	sCount := strconv.Itoa(count)
	ret = append(ret, []byte(sCount)...)
	ret = append(ret, last)

	return ret
}
