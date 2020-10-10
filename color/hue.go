package color

import (
	"errors"
	"fmt"
	"math"
)

var ErrInvalidHue = errors.New("invalid hue value, must be in the range of [0, 360)")

func NewHue(val float64) Hue {
	return Hue{Val: normalizeHueValue(val)}
}

// Represents a color wheel in degrees
// Should be created with NewHue()
type Hue struct {
	Val float64 // a value that must be in the range of [0, 360)
}

func (hue Hue) IsValid() bool {
	return hue.Val >= 0 && hue.Val < 360
}

func (hue Hue) String() string {
	return fmt.Sprintf("%.1f\u00B0", hue.Val)
}

// If delta is positive, it moves the hue clockwise
// If delta is negative, it moves the hue counter clockwise
func MoveHue(hue Hue, delta float64) Hue {
	if !hue.IsValid() {
		panic(ErrInvalidHue)
	}
	return Hue{
		Val: normalizeHueValue(hue.Val + delta),
	}
}

func HueDistanceCW(from Hue, to Hue) float64 {
	if from.Val == to.Val {
		return 0
	} else if from.Val < to.Val {
		return to.Val - from.Val
	} else {
		return (to.Val + 360) - from.Val
	}
}

func HueDistanceCCW(from Hue, to Hue) float64 {
	if from.Val == to.Val {
		return 0
	} else if from.Val > to.Val {
		return to.Val - from.Val
	} else {
		return (to.Val - 360) - from.Val
	}
}

func HueDistanceNearest(from Hue, to Hue) float64 {
	cwd := HueDistanceCW(from, to)
	ccwd := HueDistanceCCW(from, to)

	acwd := math.Abs(cwd)
	accwd := math.Abs(ccwd)

	if acwd > accwd {
		return ccwd
	} else {
		return cwd
	}
}

func normalizeHueValue(val float64) float64 {
	if val >= 0 && val < 360 {
		return val
	} else if val >= 360 {
		rev := math.Floor(val / 360)
		val -= 360 * rev
	} else if val < 0 {
		rev := math.Floor(math.Abs(val / 360))
		val += 360 + (360 * rev)
	}
	return val
}

func RGBtoHue(rgb RGB) Hue {
	min := math.Min(math.Min(rgb.R, rgb.G), rgb.B)
	max := math.Max(math.Max(rgb.R, rgb.G), rgb.B)
	c := max - min

	var hP float64
	switch max {
	case rgb.R:
		hP = math.Mod((rgb.G-rgb.B)/c, 6)
	case rgb.G:
		hP = ((rgb.B - rgb.R) / c) + 2
	case rgb.B:
		hP = ((rgb.R - rgb.G) / c) + 4
	}

	h := hP * 60

	return NewHue(h)
}

func computeRGB(c, x, hP, m float64) RGB {
	var r1, g1, b1 float64

	switch {
	case 0 <= hP && hP <= 1:
		r1, g1, b1 = c, x, 0
	case 1 <= hP && hP <= 2:
		r1, g1, b1 = x, c, 0
	case 2 <= hP && hP <= 3:
		r1, g1, b1 = 0, c, x
	case 3 <= hP && hP <= 4:
		r1, g1, b1 = 0, x, c
	case 4 <= hP && hP <= 5:
		r1, g1, b1 = x, 0, c
	case 5 <= hP && hP <= 6:
		r1, g1, b1 = c, 0, x
	default:
		r1, g1, b1 = 0, 0, 0
	}

	return RGB{
		R: r1 + m,
		G: g1 + m,
		B: b1 + m,
	}
}
