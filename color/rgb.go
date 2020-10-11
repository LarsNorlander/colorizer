package color

import "fmt"

type RGB struct {
	R, G, B float64
}

func (rgb RGB) FormalString() string {
	return fmt.Sprintf("rgb(%f,%f,%f)", rgb.R, rgb.G, rgb.B)
}

func (rgb RGB) String() string {
	return rgb.ToHex().String()
}

func BlendRGB(x RGB, y RGB) RGB {
	return PartialBlendRGB(x, y, 0.5)
}

func PartialBlendRGB(x RGB, y RGB, percentage float64) RGB {
	return RGB{
		R: wavg(x.R, y.R, percentage),
		G: wavg(x.G, y.G, percentage),
		B: wavg(x.B, y.B, percentage),
	}
}

func RGBGradient(between int, rgb ...RGB) []RGB {
	grad := make([]RGB, len(rgb)+(between*(len(rgb)-1)))

	steps := float64(between) + 1
	weight := 1.0 / steps

	grad[0] = rgb[0]
	for i := 0; i < len(rgb)-1; i++ {
		ca := rgb[i]
		cb := rgb[i+1]
		curWeight := 0.0
		offset := i * (between + 1)
		for j := 0; j < between+2; j++ {
			grad[j+offset] = PartialBlendRGB(ca, cb, curWeight)
			curWeight += weight
		}
	}

	return grad
}
