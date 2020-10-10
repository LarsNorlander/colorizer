package color

import (
	"fmt"
)

type RGB struct {
	R, G, B float64
}

func (rgb RGB) String() string {
	return fmt.Sprintf("rgb(%f,%f,%f)", rgb.R, rgb.G, rgb.B)
}

func Blend(x RGB, y RGB) RGB {
	return RGB{
		R: avg(x.R, y.R),
		G: avg(x.G, y.G),
		B: avg(x.B, y.B),
	}
}

func RGBGradient(between int, rgb ...RGB) []RGB {
	rgbLen := len(rgb)
	grad := make([]RGB, rgbLen+(between*(rgbLen-1)))
	stepCount := between + 1

	grad[len(grad)-1] = rgb[rgbLen-1]

	for i := 0; i < rgbLen-1; i++ {
		x := rgb[i]
		y := rgb[i+1]

		rStep := calcStep(x.R, y.R, stepCount)
		gStep := calcStep(x.G, y.G, stepCount)
		bStep := calcStep(x.B, y.B, stepCount)

		rCur := x.R
		gCur := x.G
		bCur := x.B

		for j := 0; j < stepCount; j++ {
			offset := i * stepCount
			grad[j+offset] = RGB{rCur, gCur, bCur}
			rCur += rStep
			gCur += gStep
			bCur += bStep
		}
	}

	return grad
}
