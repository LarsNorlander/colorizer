package color

import (
	"fmt"
	"math"
)

type point struct {
	x, y float64
}

func (p point) String() string {
	return fmt.Sprintf("(%f, %f)", p.x, p.y)
}

type line struct {
	slope, yIntercept float64
}

func (l line) String() string {
	if l.yIntercept < 0 {
		return fmt.Sprintf("y = %fx - %f", l.slope, math.Abs(l.yIntercept))
	} else {
		return fmt.Sprintf("y = %fx + %f", l.slope, l.yIntercept)
	}
}

func wavg(x, y, w float64) float64 {
	xWeight := 1 - w
	yWeight := 1 - xWeight
	return (x * xWeight) + (y*yWeight)/1
}

func slope(a, b point) float64 {
	return (b.y - a.y) / (b.x - a.x)
}

func intersect(a, b line) point {
	slope := a.slope + (-1 * b.slope)
	intercept := (-1 * a.yIntercept) + b.yIntercept
	x := intercept / slope
	y := a.slope*x + a.yIntercept
	return point{
		x: x,
		y: y,
	}
}

func lineFromPoints(a, b point) line {
	slope := slope(a, b)
	Yi := a.y + (-1*slope)*a.x
	return line{
		slope:      slope,
		yIntercept: Yi,
	}
}

func distanceBetweenPoints(a, b point) float64 {
	return math.Sqrt(math.Pow(b.x-a.x, 2) + math.Pow(b.y-a.y, 2))
}

func areaEqTriangle(side float64) float64 {
	return (math.Sqrt(3) / 4) * math.Pow(side, 2)
}

func triangleHeight(base float64, area float64) float64 {
	return 2 * (area / base)
}
