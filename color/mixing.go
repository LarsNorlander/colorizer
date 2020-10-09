package color

func Blend(x RGB, y RGB) RGB {
	return RGB{
		R: average(x.R, y.R),
		G: average(x.G, y.G),
		B: average(x.B, y.B),
	}
}

func average(x float64, y float64) float64 {
	return (x + y) / 2
}

// TODO Implement so that it could take variadic number of colors
func GenerateRGBGradient2(x RGB, y RGB, between uint) []RGB {
	grad := make([]RGB, 2+between)
	stepCount := between + 1

	rStep := computeStep(x.R, y.R, stepCount)
	gStep := computeStep(x.G, y.G, stepCount)
	bStep := computeStep(x.B, y.B, stepCount)

	rCur := x.R
	gCur := x.G
	bCur := x.B

	for i := 0; i < len(grad); i++ {
		grad[i] = RGB{rCur, gCur, bCur}
		rCur += rStep
		gCur += gStep
		bCur += bStep
	}

	return grad
}

// Hackiest shit late night trying to get things right
func GenerateRGBGradient3(x RGB, y RGB, z RGB, between uint) []RGB {
	grad := make([]RGB, 3+(between*2))
	stepCount := between + 1

	rStep := computeStep(x.R, y.R, stepCount)
	gStep := computeStep(x.G, y.G, stepCount)
	bStep := computeStep(x.B, y.B, stepCount)

	rCur := x.R
	gCur := x.G
	bCur := x.B

	for i := 0; i < (2 + int(between)); i++ {
		grad[i] = RGB{rCur, gCur, bCur}
		rCur += rStep
		gCur += gStep
		bCur += bStep
	}

	rCur -= rStep
	gCur -= gStep
	bCur -= bStep

	rStep = computeStep(y.R, z.R, stepCount)
	gStep = computeStep(y.G, z.G, stepCount)
	bStep = computeStep(y.B, z.B, stepCount)

	for i := 1 + int(between); i < len(grad); i++ {
		grad[i] = RGB{rCur, gCur, bCur}
		rCur += rStep
		gCur += gStep
		bCur += bStep
	}

	return grad
}

func GenerateHSLGradient(x HSL, y HSL, between uint) []HSL {
	grad := make([]HSL, 2+between)
	stepCount := between + 1

	hStep := computeStep(x.H, y.H, stepCount)
	sStep := computeStep(x.S, y.S, stepCount)
	lStep := computeStep(x.L, y.L, stepCount)

	hCur := x.H
	sCur := x.S
	lCur := x.L

	for i := 0; i < len(grad); i++ {
		grad[i] = HSL{hCur, sCur, lCur}
		hCur += hStep
		sCur += sStep
		lCur += lStep
	}

	return grad
}

func GenerateBlackToWhiteGradient(x HSL, between uint) []HSL {
	grad := make([]HSL, 2+between)
	stepCount := between + 1

	lStep := computeStep(0, 1, stepCount)

	lCur := 0.0

	for i := 0; i < len(grad); i++ {
		grad[i] = HSL{x.H, x.S, lCur}
		lCur += lStep
	}

	return grad
}

func GenerateLightnessGradient(h, s, darkClip, lightClip float64, between uint) []HSL {
	grad := make([]HSL, 2+between)
	stepCount := between + 1

	lStep := computeStep(0+darkClip, 1-lightClip, stepCount)

	lCur := 0.0 + darkClip

	for i := 0; i < len(grad); i++ {
		grad[i] = HSL{h, s, lCur}
		lCur += lStep
	}

	return grad
}

func computeStep(x float64, y float64, steps uint) float64 {
	diff := y - x
	return diff / float64(steps)
}
