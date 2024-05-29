package clockface

import (
	"math"
	"time"
)

// A Point represents a two-dimensional Cartesian coordinate
type Point struct {
	X float64
	Y float64
}

// SecondHand is the unit vector of the second hand of an analogue clock at time `t`
// represented as a Point.
func SecondHand(t time.Time) Point {
	p := secondHandPoint(t)
	p = Point{p.X * 90, p.Y * 90}   // scale: the length of our second hand is 90. This is just a convention.
	p = Point{p.X, -p.Y}            // flip: to account for the SVG having an origin in the top left hand corner.
	p = Point{p.X + 150, p.Y + 150} // translate: the origin needs to be the center of the clock at 150, 150!
	return p
}
func secondsInRadians(t time.Time) float64 {
	return (math.Pi / (30 / (float64(t.Second()))))
}

func secondHandPoint(t time.Time) Point {
	angle := secondsInRadians(t)
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}