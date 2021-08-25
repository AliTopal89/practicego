package clockface

import (
	"fmt"
	"io"
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
const minuteHandLength = 80

//SVGWriter writes an SVG representation of an analogue clock, showing the time t, to the writer w
func SVGWriter(w io.Writer, t time.Time) {
	io.WriteString(w, svgStart)
	io.WriteString(w, bezel)
	secondHandSVG(w, t)
	minuteHand(w, t)
	io.WriteString(w, svgEnd)
}

func secondHandSVG(w io.Writer, t time.Time) {
	p := makeHand(secondHandPoint(t), secondHandLength)
	// p = Point{p.X * secondHandLength, p.Y * secondHandLength}
	// p = Point{p.X, -p.Y}
	// p = Point{p.X + clockCentreX, p.Y + clockCentreY}
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
}

func minuteHand(w io.Writer, t time.Time) {
	p := makeHand(minuteHandPoint(t), minuteHandLength)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:3px;"/>`, p.X, p.Y)
}

func makeHand(p Point, length float64) Point {
	// scale, it to the length of hand
	// flip, it over X axis to account for the
	p = Point{p.X * length, p.Y * length}
	p = Point{p.X, -p.Y}
	// SVG having an origin in the top left hand corner
	return Point{p.X + clockCentreX, p.Y + clockCentreY}
}

// SecondHand is the unit vector of the second hand of an analogue clock at time `t`
// represented as a Point.
func SecondHand(t time.Time) Point {
	p := secondHandPoint(t)
	p = Point{p.X * secondHandLength, p.Y * secondHandLength} // scale, it to the length of hand
	p = Point{p.X, -p.Y}                                      // flip, it over X axis to account for the
	// SVG having an origin in the top left hand corner
	p = Point{p.X + clockCentreX, p.Y + clockCentreY} // translate, it the right position
	// (so that its coming from origin 150,150)
	return p
}

const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`

const bezel = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`

const svgEnd = `</svg>`

func secondsInRadians(t time.Time) float64 {
	return (math.Pi / (30 / (float64(t.Second()))))
}

/*
Going to use SVG of a clock for us to play with.
SVGs are a fantastic image format to manipulate programmatically because
they're written as a series of shapes, described in XML
*/

func secondHandPoint(t time.Time) Point {
	return angleToPoint(secondsInRadians(t))
}

func minutesInRadians(t time.Time) float64 {
	return (secondsInRadians(t) / 60) +
		(math.Pi / (30 / float64(t.Minute())))
}

func angleToPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}

func minuteHandPoint(t time.Time) Point {
	return angleToPoint(minutesInRadians(t))
}
