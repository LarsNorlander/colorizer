package color

import (
	"fmt"
	"math"
)

type HSL struct {
	H    Hue
	S, L float64
}

func (hsl HSL) String() string {
	return fmt.Sprintf("hsl(%f, %f, %f)", hsl.H, hsl.S, hsl.L)
}

func (rgb RGB) ToHSL() HSL {
	min := math.Min(math.Min(rgb.R, rgb.G), rgb.B)
	max := math.Max(math.Max(rgb.R, rgb.G), rgb.B)
	c := max - min
	l := (max + min) / 2

	if c == 0 {
		return HSL{
			H: NewHue(0.0),
			S: 0.0,
			L: l,
		}
	}

	h := rgb.ToHue()

	var s float64
	if l == 0 || l == 1 {
		s = 0
	} else {
		s = c / (1 - math.Abs(2*l-1))
	}

	return HSL{H: h, S: s, L: l}
}

func (hsl HSL) ToRGB() RGB {
	c := (1 - math.Abs(2*hsl.L-1)) * hsl.S
	hP := hsl.H.Val / 60
	x := c * (1 - math.Abs(math.Mod(hP, 2)-1))
	m := hsl.L - c/2
	return computeRGB(c, x, hP, m)
}

func HSLGradient(between int, hsl ...HSL) []HSL {
	return hslGradientHelper(
		HueDistanceCW,
		between,
		hsl...,
	)
}

func ReverseHSLGradient(between int, hsl ...HSL) []HSL {
	return hslGradientHelper(
		HueDistanceCCW,
		between,
		hsl...,
	)
}

func NearestHSLGradient(between int, hsl ...HSL) []HSL {
	return hslGradientHelper(
		HueDistanceNearest,
		between,
		hsl...,
	)
}

func hslGradientHelper(
	hueStepCalc func(from Hue, to Hue) float64,
	between int, hsl ...HSL,
) []HSL {
	hslLen := len(hsl)
	grad := make([]HSL, hslLen+(between*(hslLen-1)))
	stepCount := between + 1

	grad[len(grad)-1] = hsl[hslLen-1]

	for i := 0; i < hslLen-1; i++ {
		x := hsl[i]
		y := hsl[i+1]

		hStep := hueStepCalc(x.H, y.H) / float64(stepCount)
		sStep := calcStep(x.S, y.S, stepCount)
		lStep := calcStep(x.L, y.L, stepCount)

		hCur := x.H
		sCur := x.S
		lCur := x.L

		for j := 0; j < stepCount; j++ {
			offset := i * stepCount
			grad[j+offset] = HSL{hCur, sCur, lCur}
			hCur = MoveHue(hCur, hStep)
			sCur += sStep
			lCur += lStep
		}
	}

	return grad
}

func LightnessGradient(h Hue, s float64, between int, darkClip, lightClip float64) []HSL {
	grad := make([]HSL, 2+between)
	stepCount := between + 1

	lStep := calcStep(0+darkClip, 1-lightClip, stepCount)

	lCur := 0.0 + darkClip

	for i := 0; i < len(grad); i++ {
		grad[i] = HSL{h, s, lCur}
		lCur += lStep
	}

	return grad
}
