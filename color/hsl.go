package color

import (
	"fmt"
	"math"
)

type HSL struct {
	H    Hue
	S, L float64
}

func (hsl HSL) FormalString() string {
	return fmt.Sprintf("hsl(%.1f\u00B0, %.1f%%, %.1f%%)", hsl.H.Val, hsl.S * 100, hsl.L * 100)
}

func (hsl HSL) String() string {
	return hsl.ToRGB().ToHex().String()
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

func BlendHSL(x HSL, y HSL, strategy HueDistanceSolver) HSL {
	return PartialBlendHSL(x, y, 0.5, strategy)
}

func PartialBlendHSL(x HSL, y HSL, percentage float64, strategy HueDistanceSolver) HSL {
	distance := strategy(x.H, y.H)
	movement := distance * percentage

	return HSL{
		H: MoveHue(x.H, movement),
		S: wavg(x.S, y.S, percentage),
		L: wavg(x.L, y.L, percentage),
	}
}

func HSLGradient(between int, strategy HueDistanceSolver, hsl ...HSL) []HSL {
	grad := make([]HSL, len(hsl)+(between*(len(hsl)-1)))

	steps := float64(between) + 1
	weight := 1.0 / steps

	grad[0] = hsl[0]
	for i := 0; i < len(hsl)-1; i++ {
		ca := hsl[i]
		cb := hsl[i+1]
		curWeight := 0.0
		offset := i * (between + 1)
		for j := 0; j < between+2; j++ {
			grad[j+offset] = PartialBlendHSL(ca, cb, curWeight, strategy)
			curWeight += weight
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
