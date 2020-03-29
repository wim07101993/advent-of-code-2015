package day09

import "strconv"

type Distance int

func (d *Distance) String() string {
	return strconv.Itoa(int(*d))
}
