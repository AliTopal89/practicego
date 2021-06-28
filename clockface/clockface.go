package clockface

import (
	"math"
	"time"
)

// A Point represents a two dimensional Cartesian coordinate.
type Point struct {
	X float64
	Y float64
}

// SecondHand is the unit vector of the second hand of an analogue clock at time `t`.
// represented as a Point.
func SecondHand(t time.Time) Point {
	return Point{150, 60}
}

func secondsInRadians(t time.Time) float64 {
	return (math.Pi / (30 / (float64(t.Second()))))
}

/*
Going to use SVG of a clock for us to play with.
SVGs are a fantastic image format to manipulate programmatically because
they're written as a series of shapes, described in XML
*/

func SecondHandPoint(t time.Time) Point {
	return Point{
		X: 0,
		Y: -1,
	}
}
