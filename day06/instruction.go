package day06

import (
	"errors"
	"regexp"
	"strconv"
)

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

func (i *Instruction) ExecuteOn(g Grid) {
	switch i.Do {
	case Off:
		g.DoForRange(i.From, i.To, func(x, y int) {
			g.TurnOff(x, y)
		})
	case On:
		g.DoForRange(i.From, i.To, func(x, y int) {
			g.TurnOn(x, y)
		})
	case Toggle:
		g.DoForRange(i.From, i.To, func(x, y int) {
			g.Toggle(x, y)
		})
	}
}

