package clockface

import (
	"fmt"
	"math"
	"time"
)

// A Point represents a two dimensional Cartesian coordinate.
type Point struct {
	X float64
	Y float64
}

const secondHandLength = 90
const clockCentreX = 150
const clockCentreY = 150

// SecondHand is the unit vector of the second hand of an analogue clock at time `t`
// represented as a Point.
func SecondHand(t time.Time) Point {
	p := SecondHandPoint(t)
	p = Point{p.X * secondHandLength, p.Y * secondHandLength} // scale, it to the length of hand
	p = Point{p.X, -p.Y}                                      // flip, it over X axis to account for the
	// SVG having an origin in the top left hand corner
	p = Point{p.X + clockCentreX, p.Y + clockCentreY} // translate, it the right position
	// (so that its coming from origin 150,150)
	return p
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
	angle := secondsInRadians(t)
	X := math.Sin(angle)
	Y := math.Cos(angle)
	fmt.Println("MY preciseness", Point{X, Y})
	return Point{X, Y}
}
