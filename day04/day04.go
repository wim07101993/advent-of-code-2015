package day04

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"sync"
)

type SolverDay4 struct {
	input              []byte
	solution1          int
	solution2          int
	numberOfGoRoutines int
}

func New(inputProvider func() []byte, numberOfGoRoutines int) *SolverDay4 {
	if numberOfGoRoutines == 0 {
		panic("Day 4 will not solve without go routines")
	}
	return &SolverDay4{
		input:              inputProvider(),
		numberOfGoRoutines: numberOfGoRoutines,
	}
}

func (s *SolverDay4) SolvePart1() string {
	return strconv.Itoa(s.findHash())
}

func (s *SolverDay4) SolvePart2() string {
	return ""
}

func (s *SolverDay4) findHash() int {
	c := make(chan int)
	solution := make(chan int)
	stop := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go feedNumbers(c, stop)

	for i := 0; i < s.numberOfGoRoutines; i++ {
		go s.checkNumbers(c, solution, &wg)
	}

	wg.Wait()
	stop <- 1
	x := <-solution
	return x
}

func (s *SolverDay4) checkNumbers(c <-chan int, solution chan int, wg *sync.WaitGroup) {
	var x int
	for {
		select {
		case x = <-c:

			sx := strconv.Itoa(x)
			input := append(s.input, sx...)
			h := hash(input)
			if checkHash(h) {
				wg.Done()
				solution <- x
				return
			}

			if x%10000000 == 0 {
				fmt.Print(x / 1000000)
			} else if x%1000000 == 0 {
				fmt.Print(".")
			}
		}
	}
}

func feedNumbers(c chan<- int, stop <-chan int) {
	x := 0
	for {
		select {
		case c <- x:
			x++
		case <-stop:
			close(c)
			return
		}
	}
}

func checkHash(h []byte) bool {
	return h[0] == 0 && h[1] == 0 && h[2] < 16
}

func hash(input []byte) []byte {
	h := md5.New()
	h.Write(input)

	return h.Sum(nil)
}
