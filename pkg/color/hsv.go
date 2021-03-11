package color

import (
	"fmt"
	. "math"
)

type HSV struct {
	H    Hue
	S, V float64
}

func (hsv HSV) String() string {
	return fmt.Sprintf("hsv(%s, %.1f%%, %.1f%%)", hsv.H, hsv.S, hsv.V)
}

func (rgb RGB) HSV() HSV {
	min := Min(Min(rgb.R, rgb.G), rgb.B)
	max := Max(Max(rgb.R, rgb.G), rgb.B)
	c := max - min
	v := max

	h := rgb.Hue()

	var s float64
	if v == 0 {
		s = 0
	} else {
		s = c / v
	}

	return HSV{H: h, S: s, V: v}
}

func (hsv HSV) RGB() RGB {
	c := hsv.V * hsv.S
	hP := hsv.H.Val / 60
	x := c * (1 - Abs(Mod(hP, 2)-1))
	m := hsv.V - c
	return computeRGB(c, x, hP, m)
}

func PartialHSVBlend(a Color, b Color, percentage float64, strategy HueDistanceSolver) Color {
	x := a.RGB().HSV()
	y := b.RGB().HSV()

	distance := strategy(x.H, y.H)
	movement := distance * percentage

	return HSV{
		H: MoveHue(x.H, movement),
		S: wavg(x.S, y.S, percentage),
		V: wavg(x.V, y.V, percentage),
	}
}

func HSVGradient(between int, strategy HueDistanceSolver, colors ...Color) []Color {
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
			grad[j+offset] = PartialHSVBlend(ca, cb, curWeight, strategy)
			curWeight += weight
		}
	}

	return grad
}
