package color

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHue(t *testing.T) {
	tests := []struct {
		name     string
		value    float64
		expected Hue
	}{
		{
			name:     "create hue in range",
			value:    55,
			expected: Hue{55},
		},
		{
			name:     "create hue at 0",
			value:    0,
			expected: Hue{0},
		},
		{
			name:     "create hue at 360",
			value:    360,
			expected: Hue{0},
		},
		{
			name:     "create hue with 1 revolution",
			value:    365,
			expected: Hue{5},
		},
		{
			name:     "create hue using negative value",
			value:    -5,
			expected: Hue{355},
		},
		{
			name:     "create hue using negative value with 2 revolutions",
			value:    -725,
			expected: Hue{355},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, NewHue(test.value))
		})
	}
}

func TestMoveHue(t *testing.T) {
	tests := []struct {
		name     string
		start    Hue
		delta    float64
		expected Hue
	}{
		{
			name:     "move clockwise",
			start:    Hue{3},
			delta:    40,
			expected: Hue{43},
		},
		{
			name:     "move counter clockwise",
			start:    Hue{40},
			delta:    -3,
			expected: Hue{37},
		},
		{
			name:     "move clockwise beyond 360",
			start:    Hue{355},
			delta:    10,
			expected: Hue{5},
		},
		{
			name:     "move counter clockwise beyond 0",
			start:    Hue{5},
			delta:    -10,
			expected: Hue{355},
		},
		{
			name:     "move clockwise to 359.5 deg",
			start:    Hue{350},
			delta:    9.5,
			expected: Hue{359.5},
		},
		{
			name:     "move counter clockwise to 0.5 deg",
			start:    Hue{5},
			delta:    -4.5,
			expected: Hue{0.5},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := MoveHue(test.start, test.delta)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestMoveHuePanicConditions(t *testing.T) {
	tests := []struct {
		name  string
		hue   Hue
		panic bool
	}{
		{
			name:  "panic at hue 360",
			hue:   Hue{360},
			panic: true,
		},
		{
			name:  "panic at hue -0.5",
			hue:   Hue{-0.5},
			panic: true,
		},
		{
			name:  "don't panic at hue 0",
			hue:   Hue{0},
			panic: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil && test.panic == true {
					t.Errorf("should have paniced at hue value of %f", test.hue.Val)
				}
			}()

			MoveHue(test.hue, 0)
		})
	}
}

func TestHueDistanceCW(t *testing.T) {
	tests := []struct {
		name     string
		from     Hue
		to       Hue
		expected float64
	}{
		{
			name:     "from 10 degrees to 10 degrees",
			from:     Hue{10},
			to:       Hue{10},
			expected: 0,
		},
		{
			name:     "from 3 degrees to 10 degrees",
			from:     Hue{3},
			to:       Hue{10},
			expected: 7,
		},
		{
			name:     "from 10 degrees to 3 degrees",
			from:     Hue{10},
			to:       Hue{3},
			expected: 353,
		},
		{
			name:     "180 degrees apart",
			from:     Hue{90},
			to:       Hue{270},
			expected: 180,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, HueDistanceCW(test.from, test.to))
		})
	}
}

func TestHueDistanceCCW(t *testing.T) {
	tests := []struct {
		name     string
		from     Hue
		to       Hue
		expected float64
	}{
		{
			name:     "from 10 degrees to 10 degrees",
			from:     Hue{10},
			to:       Hue{10},
			expected: 0,
		},
		{
			name:     "from 3 degrees to 10 degrees",
			from:     Hue{3},
			to:       Hue{10},
			expected: -353,
		},
		{
			name:     "from 10 degrees to 3 degrees",
			from:     Hue{10},
			to:       Hue{3},
			expected: -7,
		},
		{
			name:     "180 degrees apart",
			from:     Hue{90},
			to:       Hue{270},
			expected: -180,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, HueDistanceCCW(test.from, test.to))
		})
	}
}

func TestHueDistanceNearest(t *testing.T) {
	tests := []struct {
		name     string
		from     Hue
		to       Hue
		expected float64
	}{
		{
			name:     "from 10 degrees to 10 degrees",
			from:     Hue{10},
			to:       Hue{10},
			expected: 0,
		},
		{
			name:     "from 3 degrees to 10 degrees",
			from:     Hue{3},
			to:       Hue{10},
			expected: 7,
		},
		{
			name:     "from 10 degrees to 3 degrees",
			from:     Hue{10},
			to:       Hue{3},
			expected: -7,
		},
		{
			name:     "equal distance should prefer clockwise",
			from:     Hue{90},
			to:       Hue{270},
			expected: 180,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, HueDistanceNearest(test.from, test.to))
		})
	}
}
