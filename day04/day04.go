package day04

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

type SolverDay4 struct {
	input     []byte
	solution1 int
	solution2 int
}

func New(inputProvider func() []byte) *SolverDay4 {
	return &SolverDay4{
		input: inputProvider(),
	}
}

func (s *SolverDay4) SolvePart1() string {
	return strconv.Itoa(s.findHash(5))
}

func (s *SolverDay4) SolvePart2() string {
	return strconv.Itoa(s.findHash(6))
}

func (s *SolverDay4) findHash(leadingZeroes int) int {
	feedStop := make(chan int)
	num := make(chan int)
	sol := make(chan int, 10)

	go feedNumbers(num, feedStop)

	go s.checkNumbers(num, sol, leadingZeroes, func(n int) {
		if n%1000000 == 0 && n != 0 {
			fmt.Print(n / 1000000)
		} else if n%100000 == 0 && n != 0 {
			fmt.Print(".")
		}
	})

	x := -1
	for x < 0 {
		x = <-sol
	}

	feedStop <- 1

	return x
}

func (s *SolverDay4) checkNumbers(num <-chan int, solution chan<- int, leadingZeros int, callback func(int)) {
	var x int
	for {
		select {
		case x = <-num:
			go s.checkNumber(x, solution, leadingZeros, callback)
		}
	}
}

func (s *SolverDay4) checkNumber(num int, solution chan<- int, leadingZeros int, callback func(int)) {
	sx := strconv.Itoa(num)
	input := append(s.input, sx...)
	h := hash(input)
	if checkHash(h, leadingZeros) {
		solution <- num
		return
	}
	solution <- -1
	callback(num)
}

func feedNumbers(c chan<- int, stop <-chan int) {
	x := 0
	for {
		select {
		case c <- x:
			x++
		case <-stop:
			return
		}
	}
}

func checkHash(h []byte, leadingZeroes int) bool {
	ok := true

	// leadingZeroes = 5			// leadingZeroes = 6
	// r = 5%2 = 1					// r = 6%2 = 0
	r := leadingZeroes % 2
	// mi = 5/2 = 2					// mi = 6/2 = 3
	mi := leadingZeroes / 2
	// h[0] == 0; h[1]				// h[0] == 0; h[1] == 0; h[2] == 0
	for i := 0; i < mi; i++ {
		ok = ok && h[i] == 0
	}

	if r > 0 {
		// h[2] < 16
		ok = ok && h[mi] < 16
	}

	if ok {
		fmt.Printf("%x\r\n", h)
	}
	return ok
}

func hash(input []byte) []byte {
	h := md5.New()
	h.Write(input)

	return h.Sum(nil)
}
