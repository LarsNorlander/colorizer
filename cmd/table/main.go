package main

import (
	"fmt"
	. "github.com/LarsNorlander/colorizer/color"
)

func main() {
	stdcw := GenerateColorWheelFromRGB(
		MustParseHex("#FF0000"),
		MustParseHex("#00FF00"),
		MustParseHex("#0000FF"),
	)

	cw, _ := GenerateColorWheel(map[Name]Color{
		Red:         MustParseHex("#FF4D65"),
		Orange:      MustParseHex("#FFAA00"),
		Yellow:      MustParseHex("#FFDE00"),
		GreenCyan:   MustParseHex("#00FFB3"),
		Cyan:        MustParseHex("#00E1FF"),
		CyanBlue:    MustParseHex("#24b7ff"),
		Blue:        MustParseHex("#478cff"),
		BlueMagenta: MustParseHex("#8453ff"),
	})
	blk := MustParseHex("#1f2022")
	wht := MustParseHex("#f7f8fa")
	mapper := NewMapper(cw, blk, wht)

	grays := HSLGradient(10, HueDistanceCW, blk, wht)
	for i := 0; i < 7; i++ {
		fmt.Print(grays[i].RGB().Hex())
	}
	fmt.Println()
	for i := 5; i < 12; i++ {
		fmt.Print(grays[i].RGB().Hex())
	}
	fmt.Println()

	for i := 0; i < 12; i++ {
		line := RGBGradient(2,
			MustParseHex("#000000"),
			stdcw.GetAt(i),
			MustParseHex("#FFFFFF"),
		)
		for _, color := range line {
			fmt.Print(mapper(color).RGB().Hex())
		}
		fmt.Println()
	}
}
