package color

import "math"

func slope(a, b Point) float64 {
	return (b.Y - a.Y) / (b.X - a.X)
}

func intersect(a, b Line) Point {
	slope := a.Slope + (-1 * b.Slope)
	intercept := (-1 * a.Yi) + b.Yi
	x := intercept / slope
	y := a.Slope*x + a.Yi
	return Point{
		X: x,
		Y: y,
	}
}

func calculateLine(a, b Point) Line {
	slope := slope(a, b)
	Yi := a.Y + (-1*slope)*a.X
	return Line{
		Slope: slope,
		Yi:    Yi,
	}
}

func calcDistance(a, b Point) float64 {
	return math.Sqrt(math.Pow(b.X-a.X, 2) + math.Pow(b.Y-a.Y, 2))
}

type Point struct {
	X, Y float64
}

type Line struct {
	Slope, Yi float64
}
