package main

import (
	"fmt"
	. "github.com/LarsNorlander/colorizer/color"
	"log"
)

func main() {
	cw, err := GenerateColorWheel(map[Name]Color{
		Red:  MustParseHex("#f92d0f"),
		Blue: MustParseHex("#0f84f9"),
		//Orange:      MustParseHex("#FFAA00"),
		Yellow: MustParseHex("#f9f00f"),
		//GreenCyan:   MustParseHex("#00FFB3"),
		//Cyan:        MustParseHex("#00E1FF"),
		//CyanBlue:    MustParseHex("#24b7ff"),
		//BlueMagenta: MustParseHex("#8453ff"),
		Green: MustParseHex("#0ff9a1"),
	})
	if err != nil {
		log.Fatal(err)
	}
	blk := MustParseHex("#000000")
	wht := MustParseHex("#ffffff")
	acnt := MustParseHex("#5DA58A")

	gray := HSLBlend(blk, wht, HueDistanceCW)
	grays := HSLGradient(7, HueDistanceCW, blk, gray, wht)
	for _, color := range grays {
		fmt.Print(color.RGB().Hex())
	}
	fmt.Println()

	for i := 0; i < 12; i++ {
		line := RGBGradient(6,
			MustParseHex("#000000"),
			cw.GetAt(i),
			MustParseHex("#FFFFFF"),
		)
		for _, color := range line {
			fmt.Print(color.RGB().Hex())
		}
		fmt.Print("    " + NameString[Name(i)])
		fmt.Println()
	}

	fmt.Println()
	fmt.Println("Tones:") // 12 steps
	fmt.Println()

	for i := 0; i < 12; i++ {
		line := RGBGradient(12,
			gray,
			cw.GetAt(i),
		)
		for _, color := range line {
			fmt.Print(color.RGB().Hex())
		}
		fmt.Println()
	}

	fmt.Println()
	fmt.Println("Shades:") // 12 steps
	fmt.Println()

	line := RGBGradient(12,
		blk,
		acnt,
	)
	for _, color := range line {
		fmt.Print(color.RGB().Hex())
	}
	fmt.Println()

}
