package color

func GenerateColorWheel(red Color, green Color, blue Color) *ColorWheel {
	const (
		Red = iota
		Orange
		Yellow
		YellowGreen
		Green
		GreenCyan
		Cyan
		CyanBlue
		Blue
		BlueMagenta
		Magenta
		MagentaRed
	)

	cw := NewColorWheel()

	// Set RGB
	cw.Jump(Red)
	cw.Set(red)

	cw.Jump(Green)
	cw.Set(green)

	cw.Jump(Blue)
	cw.Set(blue)

	cw.Jump(Yellow)
	yellow := HSLBlend(red, green, HueDistanceCW)
	cw.Set(yellow)

	cw.Jump(Cyan)
	cyan := HSLBlend(green, blue, HueDistanceCW)
	cw.Set(cyan)

	cw.Jump(Magenta)
	magenta := HSLBlend(blue, red, HueDistanceCW)
	cw.Set(magenta)

	// Set Tertiary
	cw.Jump(Orange)
	orange := HSLBlend(red, yellow, HueDistanceCW)
	cw.Set(orange)

	cw.Jump(YellowGreen)
	greenYellow := HSLBlend(yellow, green, HueDistanceCW)
	cw.Set(greenYellow)

	cw.Jump(GreenCyan)
	greenCyan := HSLBlend(green, cyan, HueDistanceCW)
	cw.Set(greenCyan)

	cw.Jump(CyanBlue)
	cyanBlue := HSLBlend(cyan, blue, HueDistanceCW)
	cw.Set(cyanBlue)

	cw.Jump(BlueMagenta)
	blueMagenta := HSLBlend(blue, magenta, HueDistanceCW)
	cw.Set(blueMagenta)

	cw.Jump(MagentaRed)
	magentaRed := HSLBlend(magenta, red, HueDistanceCW)
	cw.Set(magentaRed)

	cw.Jump(Red) // Reset pointer

	return cw
}
