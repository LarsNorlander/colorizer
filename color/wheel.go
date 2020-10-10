package color

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

// Set a new value for the current pointer
func (cw *ColorWheel) Set(rgb RGB) {
	cw.current.value = rgb
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
	str := ""
	cur := cw.start
	for {
		str += cur.value.ToHex().String()
		str += ", "
		cur = cur.next
		if cur == cw.start {
			break
		}
	}
	str += cur.value.ToHex().String()
	return str
}

type node struct {
	previous *node
	value    RGB
	next     *node
}
