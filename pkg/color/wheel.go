package color

import (
	"math"
)

func NewColorWheel() *ColorWheel {
	start := &node{}
	cw := ColorWheel{
		start: start,
		size:  12,
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

// A circular, doubly linked list representing the color wheel
//goland:noinspection GoNameStartsWithPackageName
type ColorWheel struct {
	start *node
	size  int
}

// Get the value at a specific index
func (cw *ColorWheel) Get(index int) Color {
	if index < 0 {
		panic("index out of bounds")
	}
	index %= 12

	current := cw.start
	for i := 0; i < index; i++ {
		current = current.next
	}
	return current.value
}

// Set the color for a specific index of the color wheel
func (cw *ColorWheel) Set(index int, color Color) {
	if index < 0 {
		panic("index out of bounds")
	}
	index %= 12

	current := cw.start
	for i := 0; i < index; i++ {
		current = current.next
	}

	current.value = color
}

// Grab the RGB value at a specific point in the color wheel.
func (cw *ColorWheel) Sample(degrees float64) Color {
	degrees = normalizeDegrees(degrees)

	index := int(math.Floor(degrees / 30))
	start := cw.getNodeAt(index)

	blendPercentage := (degrees - (float64(index) * 30)) / 30

	return PartialHSLBlend(
		start.value,
		start.next.value,
		blendPercentage,
		HueDistanceCW,
	)
}

func (cw ColorWheel) String() string {
	str := "["
	cur := cw.start
	for {
		if cur.value == nil {
			str += "nil"
		} else {
			str += cur.value.RGB().Hex().String()
		}
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
	value    Color
	next     *node
}

func (cw *ColorWheel) getNodeAt(index int) *node {
	if index < 0 {
		panic("index out of bounds")
	}
	index %= 12

	current := cw.start
	for i := 0; i < index; i++ {
		current = current.next
	}
	return current
}
