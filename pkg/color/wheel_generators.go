package color

import "errors"

func GenerateColorWheel(colors map[Name]Color) (*ColorWheel, error) {
	if len(colors) < 2 {
		return nil, errors.New("at least two colors should be provided")
	}

	cw := NewColorWheel()
	// Set the colors where already provided
	for i, color := range colors {
		cw.Set(int(i), color)
	}

	// Figure out where the gaps are
	var start int
	for i, color := range colors {
		if color == nil {
			continue
		}
		start = int(i)
		break
	}

	type pair struct {
		x, y int
	}
	var gaps []pair

	counter := start + 1
	last := start
	for {
		index := counter
		color := cw.Get(index)

		if color != nil {
			gaps = append(gaps, pair{x: last, y: counter})
			last = index
		}

		if index%12 == start {
			break
		}
		counter++
	}

	// Fill in the gaps
	for _, gap := range gaps {
		gapSize := gap.y - gap.x - 1
		colors := HSLGradient(gapSize, HueDistanceCW, cw.Get(gap.x), cw.Get(gap.y))
		for i, color := range colors {
			cw.Set(gap.x+i, color)
		}
	}

	return cw, nil
}
