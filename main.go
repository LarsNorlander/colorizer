package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(blend(rgb(26, 31, 39), rgb(255, 77, 101)))
	step := steps(26, 255, 12)
	fmt.Println(step)
	val := 26.0
	for i := 0; i < 13; i++ {
		fmt.Printf("step: %02d\t%f\t%f\n", i, val, math.Round(val))
		val += step
	}
}

type Color struct {
	R uint8
	G uint8
	B uint8
}

func rgb(r uint8, g uint8, b uint8) Color {
	return Color{
		R: r,
		G: g,
		B: b,
	}
}

func (c Color) String() string {
	return c.AsHex()
}

func (c Color) AsRGB() string {
	return fmt.Sprintf("rgb(%d,%d,%d)", c.R, c.G, c.B)
}

func (c Color) AsHex() string {
	return fmt.Sprintf("#%x%x%x", c.R, c.G, c.B)
}

func blend(x Color, y Color) Color {
	return Color{
		R: average(x.R, y.R),
		G: average(x.G, y.G),
		B: average(x.B, y.B),
	}
}

func average(x uint8, y uint8) uint8 {
	sum := int(x) + int(y)
	avg := float64(sum) / 2
	ciel := math.Round(avg)
	return uint8(ciel)
}

func steps(x uint8, y uint8, steps uint) float64 {
	diff := int(y) - int(x)
	return float64(diff) / float64(steps)
}

func gradient(x Color, y Color) [15]Color {
	grad := [15]Color{}
	grad[0] = x
	grad[14] = y
	return grad
}
