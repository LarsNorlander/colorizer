package color

import (
	"fmt"
	"math"
)

type HSL struct {
	H    Hue
	S, L float64
}

func (rgb RGB) HSL() HSL {
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

	h := rgb.Hue()

	var s float64
	if l == 0 || l == 1 {
		s = 0
	} else {
		s = c / (1 - math.Abs(2*l-1))
	}

	return HSL{H: h, S: s, L: l}
}

func (hsl HSL) RGB() RGB {
	c := (1 - math.Abs(2*hsl.L-1)) * hsl.S
	hP := hsl.H.Val / 60
	x := c * (1 - math.Abs(math.Mod(hP, 2)-1))
	m := hsl.L - c/2
	return computeRGB(c, x, hP, m)
}

func (hsl HSL) String() string {
	return fmt.Sprintf("hsl(%s, %.1f%%, %.1f%%)", hsl.H, hsl.S*100, hsl.L*100)
}

func HSLBlend(a Color, b Color, strategy HueDistanceSolver) Color {
	return PartialHSLBlend(a, b, 0.5, strategy)
}

func PartialHSLBlend(a Color, b Color, percentage float64, strategy HueDistanceSolver) Color {
	x := a.RGB().HSL()
	y := b.RGB().HSL()

	distance := strategy(x.H, y.H)
	movement := distance * percentage

	return HSL{
		H: MoveHue(x.H, movement),
		S: wavg(x.S, y.S, percentage),
		L: wavg(x.L, y.L, percentage),
	}
}

func HSLGradient(between int, strategy HueDistanceSolver, colors ...Color) []Color {
	grad := make([]Color, len(colors)+(between*(len(colors)-1)))

	steps := float64(between) + 1
	weight := 1.0 / steps

	grad[0] = colors[0]
	for i := 0; i < len(colors)-1; i++ {
		ca := colors[i]
		cb := colors[i+1]
		curWeight := 0.0
		offset := i * (between + 1)
		for j := 0; j < between+2; j++ {
			grad[j+offset] = PartialHSLBlend(ca, cb, curWeight, strategy)
			curWeight += weight
		}
	}

	return grad
}

func HSLLumGradient(h Hue, s, darkClip, lightClip float64, between int) []Color {
	grad := make([]Color, 2+between)
	stepCount := between + 1

	lStep := ((0 + darkClip) - (1 - lightClip)) / float64(stepCount)

	lCur := 0.0 + darkClip

	for i := 0; i < len(grad); i++ {
		grad[i] = HSL{h, s, lCur}
		lCur += lStep
	}

	return grad
}
