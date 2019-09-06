package day07

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Action int

const (
	Move Action = iota
	And
	Or
	Not
	Lshift
	Rshift
)

type Instruction struct {
	InputWires []string
	OutputWire string
	Action     Action
}

type Program struct {
	instructions []*Instruction
	done         []*Instruction
	todo         []*Instruction
	Variables    map[string]uint16
}

type SolverDay7 struct {
	inputProvider func() []string
	input         []string
	solution1     int
	solution2     int
	wires         map[string]uint16
}

func New(inputProvider func() []string) *SolverDay7 {
	return &SolverDay7{
		inputProvider: inputProvider,
		wires:         make(map[string]uint16),
	}
}

func (s *SolverDay7) SolvePart1() string {
	if s.input == nil {
		s.input = s.inputProvider()
	}

	p := NewProgram(s.input)
	p.Execute()
	return strconv.Itoa(int(p.Variables["a"]))
}

func (s *SolverDay7) SolvePart2() string {
	if s.input == nil {
		s.input = s.inputProvider()
	}

	return ""
}

func ParseInstruction(i string) *Instruction {
	ss := strings.Split(i, " ")
	instr := &Instruction{}

	if ss[0] == "NOT" {
		// bitwise invert
		instr.Action = Not
		instr.InputWires = []string{ss[1]}
	} else if ss[1] == "->" {
		instr.Action = Move
		instr.InputWires = []string{ss[0]}
	} else {
		instr.InputWires = []string{ss[0], ss[2]}

		switch ss[1] {
		case "AND":
			instr.Action = And
		case "OR":
			instr.Action = Or
		case "LSHIFT":
			instr.Action = Lshift
		case "RSHIFT":
			instr.Action = Rshift
		}
	}

	instr.OutputWire = ss[len(ss)-1]
	return instr
}
func ParseInstructions(sis []string) []*Instruction {
	is := make([]*Instruction, len(sis), len(sis))
	for i, si := range sis {
		is[i] = ParseInstruction(si)
	}
	return is
}

func checkNVariables(vs []uint16, n int) error {
	if len(vs) < n {
		return errors.New(fmt.Sprintf("Not enough variables, need %d, got %d", n, len(vs)))
	}
	return nil
}

func FunMove(vs []uint16) (uint16, error) {
	if err := checkNVariables(vs, 1); err != nil {
		return 0, err
	}
	return vs[0], nil
}
func FunAnd(vs []uint16) (uint16, error) {
	if err := checkNVariables(vs, 2); err != nil {
		return 0, err
	}
	return vs[0] & vs[1], nil
}
func FunOr(vs []uint16) (uint16, error) {
	if err := checkNVariables(vs, 2); err != nil {
		return 0, err
	}
	return vs[0] | vs[1], nil
}
func FunNot(vs []uint16) (uint16, error) {
	if err := checkNVariables(vs, 1); err != nil {
		return 0, err
	}
	return 65535 - vs[0], nil
}
func FunLshift(vs []uint16) (uint16, error) {
	if err := checkNVariables(vs, 2); err != nil {
		return 0, err
	}
	return vs[0] << vs[1], nil
}
func FunRshift(vs []uint16) (uint16, error) {
	if err := checkNVariables(vs, 2); err != nil {
		return 0, err
	}
	return vs[0] >> vs[1], nil
}

func NewProgram(is []string) *Program {
	p := &Program{}

	if is == nil {
		p.todo = make([]*Instruction, 0)
	} else {
		p.todo = ParseInstructions(is)
	}

	l := len(p.todo)
	p.done = make([]*Instruction, l)
	p.Variables = make(map[string]uint16, l)

	return p
}

func (i *Instruction) Execute(vs []uint16) (uint16, error) {
	var fun func([]uint16) (uint16, error)

	switch i.Action {
	case Move:
		fun = FunMove
	case And:
		fun = FunAnd
	case Or:
		fun = FunOr
	case Not:
		fun = FunNot
	case Lshift:
		fun = FunLshift
	case Rshift:
		fun = FunRshift
	}

	return fun(vs)
}

func (i *Instruction) Values(vm map[string]uint16) []uint16 {
	vs := []uint16{}
	for _, w := range i.InputWires {
		if v, err := strconv.Atoi(w); err == nil {
			vs = append(vs, uint16(v))
		} else if v, ok := vm[w]; ok {
			vs = append(vs, v)
		} else {
			vs = nil
			break
		}
	}

	return vs
}

func (p *Program) Add(i *Instruction) {
	p.todo = append(p.todo, i)
}

func (p *Program) Execute() error {
	executedInstruction := false
	index := 0

	for len(p.todo) > 0 {
		i := p.todo[index]

		// if values can't be parsed, don't execute
		if vs := i.Values(p.Variables); vs != nil {
			r, err := i.Execute(vs)
			if err != nil {
				return err
			}
			p.Variables[i.OutputWire] = r
			p.moveTodoToDone(index)
			executedInstruction = true
		}

		// go to next instruction
		index++
		// if i is last instruction reset index
		if index >= len(p.todo) {
			// if no instructions have been executed, return error
			if !executedInstruction {
				return errors.New("Instructions don't make sense")
			}
			executedInstruction = false
			index = 0
		}
	}

	return nil
}

func (p *Program) moveTodoToDone(index int) {
	i := p.todo[index]
	p.done = append(p.done, i)
	x := index + 1
	p.todo = append(p.todo[:index], p.todo[x:]...)
}
