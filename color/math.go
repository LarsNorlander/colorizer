package color

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func (p Point) String() string {
	return fmt.Sprintf("(%f, %f)", p.X, p.Y)
}

type Line struct {
	Slope, YIntercept float64
}

func (l Line) String() string {
	if l.YIntercept < 0 {
		return fmt.Sprintf("y = %fx - %f", l.Slope, math.Abs(l.YIntercept))
	} else {
		return fmt.Sprintf("y = %fx + %f", l.Slope, l.YIntercept)
	}
}

func wavg(x, y, w float64) float64 {
	xWeight := 1 - w
	yWeight := 1 - xWeight
	return (x * xWeight) + (y*yWeight)/1
}

func slope(a, b Point) float64 {
	return (b.Y - a.Y) / (b.X - a.X)
}

func intersect(a, b Line) Point {
	slope := a.Slope + (-1 * b.Slope)
	intercept := (-1 * a.YIntercept) + b.YIntercept
	x := intercept / slope
	y := a.Slope*x + a.YIntercept
	return Point{
		X: x,
		Y: y,
	}
}

func lineFromPoints(a, b Point) Line {
	slope := slope(a, b)
	Yi := a.Y + (-1*slope)*a.X
	return Line{
		Slope:      slope,
		YIntercept: Yi,
	}
}

func distanceBetweenPoints(a, b Point) float64 {
	return math.Sqrt(math.Pow(b.X-a.X, 2) + math.Pow(b.Y-a.Y, 2))
}

func areaEqTriangle(side float64) float64 {
	return (math.Sqrt(3) / 4) * math.Pow(side, 2)
}

func triangleHeight(base float64, area float64) float64 {
	return 2 * (area / base)
}
