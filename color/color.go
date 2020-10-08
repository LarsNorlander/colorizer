package color

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

var (
	ErrInvalidHex = errors.New("invalid hex")
)

func New(R float64, G float64, B float64) Color {
	return Color{
		R: R,
		G: G,
		B: B,
	}
}

type Color struct {
	R float64
	G float64
	B float64
}

func (c Color) String() string {
	return fmt.Sprintf("rgb(%f%%,%f%%,%f%%)", c.R, c.G, c.B)
}

func (c Color) AsHex() string {
	r := uint8(math.Round(255 * c.R / 100))
	g := uint8(math.Round(255 * c.G / 100))
	b := uint8(math.Round(255 * c.B / 100))
	return fmt.Sprintf("#%X%X%X", r, g, b)
}

func ParseHex(hex string) (Color, error) {
	hex = strings.ToUpper(hex) // Ensure it's in the format the scanners will expect
	values := struct {
		R uint8
		G uint8
		B uint8
	}{}

	var err error
	switch len(hex) {
	case 7:
		_, err = fmt.Sscanf(hex, "#%02X%02X%02X", &values.R, &values.G, &values.B)
	case 4:
		_, err = fmt.Sscanf(hex, "#%1X%1X%1X", &values.R, &values.G, &values.B)
		values.R *= 17
		values.G *= 17
		values.B *= 17
	default:
		err = ErrInvalidHex
	}

	if err != nil {
		return Color{}, err
	}

	return Color{
		R: float64(values.R) / 255.0 * 100,
		G: float64(values.G) / 255.0 * 100,
		B: float64(values.B) / 255.0 * 100,
	}, nil
}

func MustParseHex(hex string) Color {
	color, err := ParseHex(hex)
	if err != nil {
		panic(err)
	}
	return color
}

func Blend(x Color, y Color) Color {
	return Color{
		R: average(x.R, y.R),
		G: average(x.G, y.G),
		B: average(x.B, y.B),
	}
}

func average(x float64, y float64) float64 {
	return (x + y) / 2
}

func GenerateGradient(x Color, y Color, between uint) []Color {
	grad := make([]Color, 2+between)
	stepCount := between + 1

	rStep := computeStep(x.R, y.R, stepCount)
	gStep := computeStep(x.G, y.G, stepCount)
	bStep := computeStep(x.B, y.B, stepCount)

	rCur := x.R
	gCur := x.G
	bCur := x.B

	for i := 0; i < len(grad); i++ {
		grad[i] = Color{rCur, gCur, bCur}
		rCur += rStep
		gCur += gStep
		bCur += bStep
	}

	return grad
}

func computeStep(x float64, y float64, steps uint) float64 {
	diff := y - x
	return diff / float64(steps)
}
