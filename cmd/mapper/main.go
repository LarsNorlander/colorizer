package main

import (
	"fmt"
	"github.com/LarsNorlander/colorizer/color"
)

func main() {
	cw := color.GenerateColorWheel(
		color.MustParseHex("#FF0000").ToRGB(),
		color.MustParseHex("#00ff00").ToRGB(),
		color.MustParseHex("#0000ff").ToRGB(),
	)
	blk := color.MustParseHex("#000000").ToRGB()
	wht := color.MustParseHex("#ffffff").ToRGB()

	rgb := []color.RGB{
		color.MustParseHex("#ff0000").ToRGB(),
		color.MustParseHex("#000000").ToRGB(),
		color.MustParseHex("#ffffff").ToRGB(),
		color.MustParseHex("#503030").ToRGB(),
		color.MustParseHex("#602020").ToRGB(),
		color.MustParseHex("#bf4040").ToRGB(),
		color.MustParseHex("#bf4040").ToRGB(),
		color.MustParseHex("#bfff40").ToRGB(),
	}

	for i := range rgb {
		original := rgb[i]
		mapped := color.MapToWold(cw, blk, wht, original)
		fmt.Printf("original: %s\nmapped  : %s\n", original, mapped)
		fmt.Println()
	}
}
