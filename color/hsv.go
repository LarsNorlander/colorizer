package color

import (
	"fmt"
	"math"
)

type HSV struct {
	H    Hue
	S, V float64
}

func (hsv HSV) String() string {
	return fmt.Sprintf("hsv(%f, %f, %f)", hsv.H, hsv.S, hsv.V)
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

func HSVGradient(between int, hsv ...HSV) []HSV {
	return hsvGradientHelper(
		HueDistanceCW,
		between,
		hsv...,
	)
}

func ReverseHSVGradient(between int, hsv ...HSV) []HSV {
	return hsvGradientHelper(
		HueDistanceCCW,
		between,
		hsv...,
	)
}

func NearestHSVGradient(between int, hsv ...HSV) []HSV {
	return hsvGradientHelper(
		HueDistanceNearest,
		between,
		hsv...,
	)
}

func hsvGradientHelper(
	hueStepCalc func(from Hue, to Hue) float64,
	between int,
	hsv ...HSV,
) []HSV {
	hsvLen := len(hsv)
	grad := make([]HSV, hsvLen+(between*(hsvLen-1)))
	stepCount := between + 1

	grad[len(grad)-1] = hsv[hsvLen-1]

	for i := 0; i < hsvLen-1; i++ {
		x := hsv[i]
		y := hsv[i+1]

		hStep := hueStepCalc(x.H, y.H) / float64(stepCount)
		sStep := calcStep(x.S, y.S, stepCount)
		lStep := calcStep(x.V, y.V, stepCount)

		hCur := x.H
		sCur := x.S
		lCur := x.V

		for j := 0; j < stepCount; j++ {
			offset := i * stepCount
			grad[j+offset] = HSV{hCur, sCur, lCur}
			hCur = MoveHue(hCur, hStep)
			sCur += sStep
			lCur += lStep
		}
	}

	return grad
}
