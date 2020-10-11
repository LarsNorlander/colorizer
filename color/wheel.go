package color

import (
	"math"
)

// A circular, doubly linked list representing the color wheel
//goland:noinspection GoNameStartsWithPackageName
type ColorWheel struct {
	start   *node
	current *node
	size    int
}

func NewColorWheel() *ColorWheel {
	start := &node{}
	cw := ColorWheel{
		start:   start,
		current: start,
		size:    12,
	}

	prev := start
	for i := 1; i < 12; i++ {
		curr := &node{
			previous: prev,
		}
		prev.next = curr
		prev = curr
	}
	prev.next = start
	start.previous = prev

	return &cw
}

// Moves current pointer to the next value, and returns it
func (cw *ColorWheel) Next() RGB {
	cw.current = cw.current.next
	return cw.current.value
}

// Moves current pointer to the previous value, and returns it
func (cw *ColorWheel) Previous() RGB {
	cw.current = cw.current.previous
	return cw.current.value
}

// Get the value at the current pointer
func (cw *ColorWheel) Get() RGB {
	return cw.current.value
}

// Get the value at a specific index
func (cw *ColorWheel) GetAt(index int) RGB {
	if index < 0 || index >= cw.size {
		panic("index out of bounds")
	}
	current := cw.start
	for i := 0; i < index; i++ {
		current = current.next
	}
	return current.value
}

func (cw *ColorWheel) getNodeAt(index int) *node {
	if index < 0 || index >= cw.size {
		panic("index out of bounds")
	}
	current := cw.start
	for i := 0; i < index; i++ {
		current = current.next
	}
	return current
}

// Set a new value for the current pointer
// TODO Implement set, so that it blends the colors next to it
func (cw *ColorWheel) Set(rgb RGB) {
	cw.current.value = rgb
}

// Grab the RGB value at a specific point in the color wheel.
func (cw *ColorWheel) Sample(degrees float64) RGB {
	degrees = normalizeDegrees(degrees)

	index := int(math.Floor(degrees / 30))
	start := cw.getNodeAt(index)

	blendPercentage := (degrees - (float64(index) * 30)) / 30

	return PartialBlendHSL(
		start.value.ToHSL(),
		start.next.value.ToHSL(),
		blendPercentage,
		HueDistanceCW, // In the future, look up the resolution strategy of the section
	).ToRGB()
}

// Move the pointer to the specified index and return the value in it, does not roll over
func (cw *ColorWheel) Jump(index int) RGB {
	if index < 0 || index >= cw.size {
		panic("index out of bounds")
	}
	cw.current = cw.start
	for i := 0; i < index; i++ {
		cw.current = cw.current.next
	}
	return cw.current.value
}

func (cw ColorWheel) String() string {
	str := "["
	cur := cw.start
	for {
		str += cur.value.ToHex().String()
		cur = cur.next
		if cur == cw.start {
			break
		}
		str += " "
	}
	str += "]"
	return str
}

type node struct {
	previous *node
	value    RGB
	next     *node
}
