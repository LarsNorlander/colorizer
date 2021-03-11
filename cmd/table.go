package cmd

import (
	"fmt"
	. "github.com/LarsNorlander/colorizer/pkg/color"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "table",
		Short: "Generate a table",
		RunE:  tableFn,
	})
}

func tableFn(cmd *cobra.Command, args []string) error {
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
		return err
	}
	blk := MustParseHex("#000000")
	wht := MustParseHex("#ffffff")
	acnt := MustParseHex("#5DA58A")

	gray := HSLBlend(blk, wht, HueDistanceCW)
	grays := HSLGradient(6, HueDistanceCW, blk, gray, wht)
	for _, color := range grays {
		fmt.Print(color.RGB().Hex())
	}
	fmt.Println()

	for i := 0; i < 12; i++ {
		line := RGBGradient(6,
			MustParseHex("#000000"),
			cw.Get(i),
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
			cw.Get(i),
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

	fmt.Println()
	fmt.Println("Tints:") // 12 steps
	fmt.Println()

	tintLine := RGBGradient(12,
		wht,
		acnt,
	)
	for _, color := range tintLine {
		fmt.Print(color.RGB().Hex())
	}
	fmt.Println()

	return nil
}
