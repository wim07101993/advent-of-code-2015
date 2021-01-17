package day09

import (
	"log"
	"regexp"
	"strconv"
)


type PointToPoint struct {
	p1 string
	p2 string
	distance int
}

var regex = regexp.MustCompile(
	`(?P<source>\w*) to (?P<destination>\w*) = (?P<distance>\d*)`)

func ParsePointToPoint(s string) PointToPoint {
	ms := regex.FindStringSubmatch(s)
	if len(ms) != 4 {
		log.Printf("Could not parse %s: Regex failed\r\n", s)
	}

	d, err := strconv.Atoi(ms[3])
	if err != nil {
		log.Printf("Could not parse %s: %s\r\n", s, err)
	}

	point1 := ms[1]
	point2 := ms[2]
	distance := d

	return PointToPoint{
		p1:       point1,
		p2:       point2,
		distance: distance,
	}
}

