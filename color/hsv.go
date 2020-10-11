package color

import (
	"fmt"
	"math"
)

type HSV struct {
	H    Hue
	S, V float64
}

func (hsv HSV) FormalString() string {
	return fmt.Sprintf("hsv(%f, %f, %f)", hsv.H, hsv.S, hsv.V)
}

func (hsv HSV) String() string {
	return hsv.ToRGB().ToHex().String()
}

func (rgb RGB) ToHSV() HSV {
	min := math.Min(math.Min(rgb.R, rgb.G), rgb.B)
	max := math.Max(math.Max(rgb.R, rgb.G), rgb.B)
	c := max - min
	v := max

	h := rgb.ToHue()

	var s float64
	if v == 0 {
		s = 0
	} else {
		s = c / v
	}

	return HSV{H: h, S: s, V: v}
}

func (hsv HSV) ToRGB() RGB {
	c := hsv.V * hsv.S
	hP := hsv.H.Val / 60
	x := c * (1 - math.Abs(math.Mod(hP, 2)-1))
	m := hsv.V - c
	return computeRGB(c, x, hP, m)
}

func PartialBlendHSV(x HSV, y HSV, percentage float64, strategy HueDistanceSolver) HSV {
	distance := strategy(x.H, y.H)
	movement := distance * percentage

	return HSV{
		H: MoveHue(x.H, movement),
		S: wavg(x.S, y.S, percentage),
		V: wavg(x.V, y.V, percentage),
	}
}

func HSVGradient(between int, strategy HueDistanceSolver, hsv ...HSV) []HSV {
	grad := make([]HSV, len(hsv)+(between*(len(hsv)-1)))

	steps := float64(between) + 1
	weight := 1.0 / steps

	grad[0] = hsv[0]
	for i := 0; i < len(hsv)-1; i++ {
		ca := hsv[i]
		cb := hsv[i+1]
		curWeight := 0.0
		offset := i * (between + 1)
		for j := 0; j < between+2; j++ {
			grad[j+offset] = PartialBlendHSV(ca, cb, curWeight, strategy)
			curWeight += weight
		}
	}

	return grad
}
