package color

func avg(x float64, y float64) float64 {
	return (x + y) / 2
}

func computeStep(x float64, y float64, steps int) float64 {
	delta := y - x
	return delta / float64(steps)
}
