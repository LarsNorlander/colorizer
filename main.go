package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	grad := gradient(hex("#1A1F27"), rgb(255, 77, 101), 13)
	for i := range grad {
		fmt.Print(grad[i])
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

func hex(s string) Color {
	var err error
	color := Color{}
	s = strings.ToUpper(s)

	switch len(s) {
	case 7:
		_, err = fmt.Sscanf(s, "#%02X%02X%02X", &color.R, &color.G, &color.B)
	case 4:
		_, err = fmt.Sscanf(s, "#%1X%1X%1X", &color.R, &color.G, &color.B)
		color.R *= 17
		color.G *= 17
		color.B *= 17
	default:
		err = fmt.Errorf("invalid hex")
	}

	if err != nil {
		panic(err)
	}
	return color
}

func (c Color) String() string {
	return c.AsHex()
}

func (c Color) AsRGB() string {
	return fmt.Sprintf("rgb(%d,%d,%d)", c.R, c.G, c.B)
}

func (c Color) AsHex() string {
	return fmt.Sprintf("#%X%X%X", c.R, c.G, c.B)
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

func gradient(x Color, y Color, between uint) []Color {
	grad := make([]Color, 2+between)
	stepCount := between + 1

	rStep := steps(x.R, y.R, stepCount)
	gStep := steps(x.G, y.G, stepCount)
	bStep := steps(x.B, y.B, stepCount)

	rCur := float64(x.R)
	gCur := float64(x.G)
	bCur := float64(x.B)

	for i := 0; i < len(grad); i++ {
		grad[i] = rgb(
			uint8(math.Round(rCur)),
			uint8(math.Round(gCur)),
			uint8(math.Round(bCur)),
		)
		rCur += rStep
		gCur += gStep
		bCur += bStep
	}

	return grad
}
