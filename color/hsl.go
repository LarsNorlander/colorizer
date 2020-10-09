package color

import (
	"fmt"
	"math"
)

type HSL struct {
	H, S, L float64
}

func (h HSL) String() string {
	return fmt.Sprintf("hsl(%f, %f, %f)", h.H, h.S, h.L)
}

func RGBtoHSL(c RGB) HSL {
	min := math.Min(math.Min(c.R, c.G), c.B)
	max := math.Max(math.Max(c.R, c.G), c.B)
	chroma := max - min
	lum := (max + min) / 2

	if chroma == 0 {
		return HSL{
			H: 0.0,
			S: 0.0,
			L: lum,
		}
	}

	var hueP float64

	switch max {
	case c.R:
		hueP = math.Mod((c.G-c.B)/chroma, 6)
	case c.G:
		hueP = ((c.B - c.R) / chroma) + 2
	case c.B:
		hueP = ((c.R - c.G) / chroma) + 4
	}

	hue := hueP * 60
	if hue < 0 {
		hue += 360
	}

	if min == max {
		return HSL{
			H: 0.0,
			S: 0.0,
			L: lum,
		}
	}

	var sat float64
	if lum == 0 || lum == 1 {
		sat = 0
	} else {
		sat = chroma / (1 - math.Abs(2*lum-1))
	}

	return HSL{
		H: hue,
		S: sat,
		L: lum,
	}
}

func HSLtoRGB(color HSL) RGB {
	c := (1 - math.Abs(2*color.L-1)) * color.S

	hP := color.H / 60

	x := c * (1 - math.Abs(math.Mod(hP, 2)-1))

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

	m := color.L - c/2

	return RGB{
		R: r1 + m,
		G: g1 + m,
		B: b1 + m,
	}
}
