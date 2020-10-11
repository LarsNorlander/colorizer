package color

func calcStep(x float64, y float64, steps int) float64 {
	delta := y - x
	return delta / float64(steps)
}

func wavg(x, y, w float64) float64 {
	xWeight := 1 - w
	yWeight := 1 - xWeight
	return (x * xWeight) + (y*yWeight)/1
}
