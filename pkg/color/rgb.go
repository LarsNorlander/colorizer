package color

import "fmt"

type RGB struct {
	R, G, B float64
}

func (rgb RGB) RGB() RGB {
	return rgb
}

func (rgb RGB) String() string {
	return fmt.Sprintf("rgb(%f,%f,%f)", rgb.R, rgb.G, rgb.B)
}

func PartialRGBBlend(a Color, b Color, percentage float64) Color {
	x := a.RGB()
	y := b.RGB()

	return RGB{
		R: wavg(x.R, y.R, percentage),
		G: wavg(x.G, y.G, percentage),
		B: wavg(x.B, y.B, percentage),
	}
}

func RGBGradient(between int, colors ...Color) []Color {
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
			grad[j+offset] = PartialRGBBlend(ca, cb, curWeight)
			curWeight += weight
		}
	}

	return grad
}
