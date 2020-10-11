package color

func GenerateColorWheel(r RGB, g RGB, b RGB) *ColorWheel {
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

	red := r.ToHSL()
	green := g.ToHSL()
	blue := b.ToHSL()

	cw := NewColorWheel()

	// Set RGB
	cw.Jump(Red)
	cw.Set(r)

	cw.Jump(Green)
	cw.Set(g)

	cw.Jump(Blue)
	cw.Set(b)

	cw.Jump(Yellow)
	yellow := BlendHSL(red, green, HueDistanceCW)
	cw.Set(yellow.ToRGB())

	cw.Jump(Cyan)
	cyan := BlendHSL(green, blue, HueDistanceCW)
	cw.Set(cyan.ToRGB())

	cw.Jump(Magenta)
	magenta := BlendHSL(blue, red, HueDistanceCW)
	cw.Set(magenta.ToRGB())

	// Set Tertiary
	cw.Jump(Orange)
	orange := BlendHSL(red, yellow, HueDistanceCW)
	cw.Set(orange.ToRGB())

	cw.Jump(YellowGreen)
	greenYellow := BlendHSL(yellow, green, HueDistanceCW)
	cw.Set(greenYellow.ToRGB())

	cw.Jump(GreenCyan)
	greenCyan := BlendHSL(green, cyan, HueDistanceCW)
	cw.Set(greenCyan.ToRGB())

	cw.Jump(CyanBlue)
	cyanBlue := BlendHSL(cyan, blue, HueDistanceCW)
	cw.Set(cyanBlue.ToRGB())

	cw.Jump(BlueMagenta)
	blueMagenta := BlendHSL(blue, magenta, HueDistanceCW)
	cw.Set(blueMagenta.ToRGB())

	cw.Jump(MagentaRed)
	magentaRed := BlendHSL(magenta, red, HueDistanceCW)
	cw.Set(magentaRed.ToRGB())

	cw.Jump(Red) // Reset pointer

	return cw
}
