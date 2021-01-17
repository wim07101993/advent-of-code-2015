package day06

type Grid interface {
	Intensity() int
	TurnOn(x, y int)
	TurnOff(x, y int)
	Toggle(x, y int)
	DoForRange(from, to Coordinate, action func(x, y int))
}

